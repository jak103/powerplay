package server

import (
	"embed"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/utils/log"
)

//go:embed static
var content embed.FS

func setupStaticServe(app *fiber.App) {
	// app.Use("/", filesystem.New(filesystem.Config{
	// 	Root:       http.FS(content),
	// 	PathPrefix: "static",
	// }))

	dir, _ := os.Getwd()
	log.Info("static pwd: %s", dir)

	app.Static("/", "/app/backend/internal/server/static")
}
