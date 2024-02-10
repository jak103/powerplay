package unittesting

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func FiberCtx() *fiber.Ctx {
	app := fiber.New()
	return app.AcquireCtx(&fasthttp.RequestCtx{})
}
