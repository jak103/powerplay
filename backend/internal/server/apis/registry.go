package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/utils/log"
)

var (
	routes []route
)

type route struct {
	method   string
	path     string
	authType auth.Type
	handlers []fiber.Handler
}

func RegisterHandler(method, path string, authType auth.Type, handlers ...fiber.Handler) {
	r := route{
		method:   method,
		path:     path,
		authType: authType,
		handlers: handlers,
	}

	routes = append(routes, r)
}

func SetupRoutes(app *fiber.App) {
	jwt := auth.GetMiddleware()

	for _, r := range routes {
		switch r.authType {
		case auth.NONE:
			app.Add(r.method, r.path, r.handlers...)
		case auth.JWT:
			handlers := append([]fiber.Handler{jwt}, r.handlers...)
			app.Add(r.method, r.path, handlers...)
		default:
			log.Alert("Unknown auth type in route registry")
		}
	}
}
