package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/leaguemanager/internal/auth"
	"github.com/jak103/leaguemanager/internal/utils/log"
)

var (
	routes []route
)

type route struct {
	Method   string
	Path     string
	AuthType auth.Type
	Handler  fiber.Handler
}

func RegisterHandler(method, path string, authType auth.Type, handler fiber.Handler) {
	r := route{
		Method:   method,
		Path:     path,
		AuthType: authType,
		Handler:  handler,
	}

	routes = append(routes, r)
}

func setupRoutes(app *fiber.App) {
	jwt := auth.GetMiddleware()

	for _, r := range routes {
		switch r.AuthType {
		case auth.NONE:
			app.Add(r.Method, r.Path, r.Handler)
		case auth.JWT:
			app.Add(r.Method, r.Path, jwt, r.Handler)
		default:
			log.Alert("Unknown auth type in route registry")
		}
	}
}
