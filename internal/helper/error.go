package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CheckIfErrFromDbToStatusCode(err error) (errs error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return Duplicate
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return BadRequest
	}
	return
}

func JsonErrorResponse(c *fiber.Ctx, errs error) (err error) {
	log.Debug().Err(err).Msg("err")
	StatusCode, err := strconv.Atoi(errs.Error())
	if err != nil {
		log.Info().Err(err).Msg("Json error response error logs")
		return c.Status(500).JSON(ErrorResponse{
			Message:   "failed translating error, internal server error!",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	//can be using switch case but time too short for debug
	if errors.Is(errs, ErrNotFound) {
		return c.Status(StatusCode).JSON(ErrorResponse{
			Message:   "Data not found",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	} else if errors.Is(errs, Duplicate) {
		return c.Status(StatusCode).JSON(ErrorResponse{
			Message:   "Check your request Unique field can't be duplicated",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	} else if errors.Is(errs, Unauthorized) {
		return c.Status(StatusCode).JSON(ErrorResponse{
			Message:   "Unauthorized",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	} else if errors.Is(errs, Forbidden) {
		return c.Status(StatusCode).JSON(ErrorResponse{
			Message:   "Forbidden",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	} else {
		return c.Status(StatusCode).JSON(ErrorResponse{
			Message:   "Internal Server Error please contact developer",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}
}
