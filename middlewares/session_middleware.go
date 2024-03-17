package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

func NewSessionMiddleware() *session.Store {
	return session.New(session.ConfigDefault)
}
