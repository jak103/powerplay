package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type createRequest struct {
	Username       string `json:"username"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phoneNumber"`
	BaseExperience int    `json:"baseExperience"`
}

type createResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   int    `json:"user_id"`
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/user", auth.Authenticated, getCurrentUser)
	apis.RegisterHandler(fiber.MethodPost, "/user", auth.Public, createUserAccount)

}

func getCurrentUser(c *fiber.Ctx) error {
	return nil
}

func createUserAccount(c *fiber.Ctx) error {

	// verify the request
	log := locals.Logger(c)
	creds := createRequest{}
	err := c.BodyParser(&creds)
	if err != nil {
		log.WithErr(err).Error("Failed to parse user creation request")
		return responder.BadRequest(c, "Failed to parse user creation request")
	}

	// TODO: Implement user creation and get a read ID
	id := 1

	createdUserResponse := createResponse{
		Username: creds.Username,
		Email:    creds.Email,
		UserId:   id,
	}

	return responder.OkWithData(c, createdUserResponse)

}
