package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/middleware"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"

	// Blank imports for apis to cause init functions to run
	_ "github.com/jak103/powerplay/internal/server/apis/auth"
	_ "github.com/jak103/powerplay/internal/server/apis/chat"
	_ "github.com/jak103/powerplay/internal/server/apis/notifications"
	_ "github.com/jak103/powerplay/internal/server/apis/user"
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: globalErrorHandler,
	})

	middleware.Setup(app)

	apis.SetupRoutes(app)

	setupStaticServe(app)

	app.Get("/files", showFiles) // TODO remove this after the files are all embedding correctly

	app.Listen(":9001")
}

// TODO remove this after the files are all embedding correctly
func showFiles(c *fiber.Ctx) error {
	dir, err := content.ReadDir("./static")
	if err != nil {
		log.WithErr(err).Error("Failed to read content dir")
	}

	files := make([]string, 0)
	for _, f := range dir {
		files = append(files, f.Name())
	}

	return c.JSON(files)
}

func globalErrorHandler(c *fiber.Ctx, err error) error {
	log := locals.Logger(c)

	log.WithErr(err).Error("Caught by global error handler")

	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
