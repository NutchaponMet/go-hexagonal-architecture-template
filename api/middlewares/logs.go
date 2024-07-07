package middlewares

import (
	"go-hexagonal/pkg/logs"
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//////////////////////////////////////////////////////////////////////
// log middleware
//////////////////////////////////////////////////////////////////////

func NewLoggerMiddleWare() func(*fiber.Ctx) error {
	return logger.New(logger.Config{
		Next:          nil,
		Done:          nil,
		Format:        "[ ${time} ] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:    "2006-01-02 15:04:05",
		TimeZone:      "Local",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: false,
	})
}

func NewZapLoggerMiddleWare() func(*fiber.Ctx) error {
	log := logs.GetLogger()
	return fiberzap.New(fiberzap.Config{
		Logger: log,
	})
}
