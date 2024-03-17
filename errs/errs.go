package errs

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotfoundError(msg string) error {
	return AppError{
		Code:    fiber.StatusNotFound,
		Message: msg,
	}
}

func NewUnexpectedError(msg string) error {
	return AppError{
		Code:    fiber.StatusInternalServerError,
		Message: msg,
	}
}
