package upload

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
	"mime/multipart"
	"os"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/upload", auth.Authenticated, handleUpload)
}

func handleUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return responder.BadRequest(c, "File not found")
	}

	if !hasUploadsDirectory() {
		if err = createUploadsDirectory(); err != nil {
			return responder.InternalServerError(c, "Could not create uploads directory")
		}
	}

	err = saveFile(c, file)
	if err != nil {
		return responder.InternalServerError(c, "Could not save file")
	}

	return responder.Ok(c, "File uploaded successfully")
}

func hasUploadsDirectory() bool {
	_, err := os.Stat("./uploads")
	return os.IsExist(err)
}

func createUploadsDirectory() error {
	return os.Mkdir("./uploads", os.ModePerm)
}

func saveFile(c *fiber.Ctx, file *multipart.FileHeader) error {
	filePath := fmt.Sprintf("../uploads/%s", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return err
	}
	return nil
}
