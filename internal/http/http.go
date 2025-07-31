package http

import (
	"fmt"
	"log"
	"os"
	internalValidator "technical_test-ayo-co-id/internal/validator"
	"time"

	"technical_test-ayo-co-id/internal/player"
	player_http "technical_test-ayo-co-id/internal/player/delivery"
	playerRepository "technical_test-ayo-co-id/internal/player/repository"
	playerUsecase "technical_test-ayo-co-id/internal/player/usecase"

	"technical_test-ayo-co-id/internal/score"
	score_http "technical_test-ayo-co-id/internal/score/delivery"
	scoreRepository "technical_test-ayo-co-id/internal/score/repository"
	scoreUsecase "technical_test-ayo-co-id/internal/score/usecase"

	"technical_test-ayo-co-id/internal/team"
	team_http "technical_test-ayo-co-id/internal/team/delivery"
	teamRepository "technical_test-ayo-co-id/internal/team/repository"
	teamUsecase "technical_test-ayo-co-id/internal/team/usecase"

	"technical_test-ayo-co-id/internal/match"
	match_http "technical_test-ayo-co-id/internal/match/delivery"
	matchRepository "technical_test-ayo-co-id/internal/match/repository"
	matchUsecase "technical_test-ayo-co-id/internal/match/usecase"

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
		player.Player{},
		match.Match{},
		score.Score{},
	)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency} ${body} ${resBody}\n",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//repository
	teamRepository := teamRepository.NewTeamPostgreRepository(db)
	playerRepository := playerRepository.NewPlayerRepository(db)
	matchRepository := matchRepository.NewMatchRepository(db)
	scoreRepository := scoreRepository.NewScoreRepository(db)

	//usecase
	scoreUsecase := scoreUsecase.NewScoreUsecase(matchRepository, scoreRepository)
	score_http.NewScoreHandler(app, myValidator, scoreUsecase)
	teamUsecase := teamUsecase.NewTeamUsecase(teamRepository)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository, teamUsecase)
	matchUsecase := matchUsecase.NewMatchUsecase(matchRepository, playerRepository, teamUsecase, scoreUsecase)

	//delivery
	player_http.NewPlayerHandler(app, myValidator, playerUsecase)
	team_http.NewTeamHandler(app, myValidator, teamUsecase)
	match_http.NewMatchHandler(app, myValidator, matchUsecase)

	log.Print(app.Listen(port))
}
