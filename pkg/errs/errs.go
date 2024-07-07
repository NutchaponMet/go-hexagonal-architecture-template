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
		Code:    fiber.ErrNotFound.Code,
		Message: msg,
	}
}

func NewUnexpectedError(msg string) error {
	return AppError{
		Code:    fiber.ErrInternalServerError.Code,
		Message: msg,
	}
}

func NewNotAcceptableError(msg string) error {
	return AppError{
		Code:    fiber.ErrNotAcceptable.Code,
		Message: msg,
	}
}

func NewTooManyArgumentsToFunction() error {
	return AppError{
		Code:    fiber.ErrInternalServerError.Code,
		Message: "Too many arguments to function.",
	}
}

func NewUnauthorizedError() error {
	return AppError{
		Code:    fiber.ErrUnauthorized.Code,
		Message: fiber.ErrUnauthorized.Message,
	}
}

func NewExitingDataError(msg string) error {
	return AppError{
		Code:    fiber.ErrNotAcceptable.Code,
		Message: msg,
	}
}
