package http

import (
	"fmt"
	"log"
	"os"
	"time"

	"technical_test-ayo-co-id/internal/team"
	team_http "technical_test-ayo-co-id/internal/team/delivery"
	teamRepository "technical_test-ayo-co-id/internal/team/repository"
	teamUsecase "technical_test-ayo-co-id/internal/team/usecase"
	internalValidator "technical_test-ayo-co-id/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	logs "log"
	_ "technical_test-ayo-co-id/docs"

	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var validate = validator.New()

func HttpRun(port string) {
	myValidator := &internalValidator.XValidator{
		Validator: validate,
	}
	log.Print(myValidator)
	newLogger := gormLogger.New(
		logs.New(os.Stdout, "\r\n", logs.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)
	zerolog.Info().Msg(viper.GetString("databse.name"))
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			viper.GetString("database.host"),
			viper.GetString("database.user"),
			viper.GetString("database.pass"),
			viper.GetString("database.name"),
			viper.GetString("database.port")),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	//migration
	db.AutoMigrate(
		team.Team{},
	)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency} ${body} ${resBody}\n",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//Team domain
	teamRepository := teamRepository.NewTeamPostgreRepository(db)
	teamUsecase := teamUsecase.NewTeamUsecase(teamRepository)
	team_http.NewTeamHandler(app, myValidator, teamUsecase)

	log.Print(app.Listen(port))
}
