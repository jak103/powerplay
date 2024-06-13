package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/logos", auth.Public, handleLogoUpload)
	apis.RegisterHandler(fiber.MethodGet, "/logos/:id<int>", auth.Public, handleGetLogoByID)
	// Need to be able to specify ctx.Type (the request's Content-Type) in the RegisterHandler
}

func handleLogoUpload(c *fiber.Ctx) error {
	log := locals.Logger(c)
	file, err := c.FormFile("image") // This is the key for the file in the form
	if err != nil {
		log.WithErr(err).Alert("Failed to upload logo")
		return responder.BadRequest(c, "Failed to upload logo")
	}

	fileData, err := file.Open()
	if err != nil {
		return responder.BadRequest(c, "Failed to open logo")
	}
	defer fileData.Close()

	imageBuffer := make([]byte, file.Size)
	if _, err = fileData.Read(imageBuffer); err != nil {
		return responder.BadRequest(c, "Failed to read logo")
	}

	db := db.GetSession(c)
	logo := models.Logo{
		Image: imageBuffer,
	}
	err = db.SaveLogo(&logo)
	if err != nil {
		log.WithErr(err).Alert("Failed to save logo")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}

func handleGetLogoByID(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	id := c.Params("id")
	logo, err := db.GetLogoByID(id)
	if err != nil {
		log.WithErr(err).Alert("Failed to get the logo from the database")
		return responder.InternalServerError(c)
	}

	c.Type("png")
	return c.Send(logo.Image)
}
