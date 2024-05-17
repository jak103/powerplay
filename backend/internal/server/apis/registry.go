package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
)

var routes map[string]route = make(map[string]route)

type route struct {
	method   string
	path     string
	roles    []auth.Role
	handlers []fiber.Handler
}

func RegisterHandler(method, path string, roles []auth.Role, handlers ...fiber.Handler) {
	r := route{
		method:   method,
		path:     path,
		roles:    roles,
		handlers: handlers,
	}

	if _, ok := routes[path]; ok {
		log.Error("Tried to register duplicate route: %s %s", method, path) // This should only happening while developing
	}

	routes[method+path] = r
}

func SetupRoutes(app *fiber.App) {
	group := app.Group("/api/v1")

	for _, r := range routes {
		group.Add(r.method, r.path, r.handlers...)
		log.Error("Tried to register duplicate route: %s %s", r.method, r.path)
	}
}

func GetRole(method, path string) []auth.Role {
	r, ok := routes[method+path]
	if !ok {
		log.Error("Trying to access an unregistered route. This should never happen")
	}

	return r.roles
}
