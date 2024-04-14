package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/users/me", auth.Authenticated, getCurrentUser)
	apis.RegisterHandler(fiber.MethodPost, "/users", auth.Public, createUserAccount)
	apis.RegisterHandler(fiber.MethodPut, "/users/:id<int>", auth.Authenticated, completeUserProfile)
	apis.RegisterHandler(fiber.MethodGet, "/users/:id<int>", auth.Authenticated, getUserById)
}

func getCurrentUser(c *fiber.Ctx) error {
	return responder.NotYetImplemented(c)
}

/*
Account creation goes in three steps
 1. Create account with email and password
 2. Verify email by clicking link
 3. Complete account by completing profile
*/
func createUserAccount(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := newSession(c)

	creds := &models.Credentials{}
	err := c.BodyParser(creds)
	if err != nil {
		log.WithErr(err).Error("Failed to parse body")
		return responder.BadRequest(c, "Failed to parse new account credentials")
	}

	if len(creds.Email) == 0 {
		return responder.BadRequest(c, "Missing email")
	}

	if len(creds.Password) == 0 {
		return responder.BadRequest(c, "Missing password")
	}

	user, err := db.GetUserByEmail(creds.Email)
	if err != nil {
		log.WithErr(err).Error("Failed to get user with email %s", creds.Email)
		return responder.InternalServerError(c)
	}

	if user != nil {
		return responder.BadRequest(c, "An account with that email address already exists")
	}

	hash, salt, err := auth.HashPassword(creds.Password, "")
	if err != nil {
		log.WithErr(err).Error("Failed to hash password")
		return responder.InternalServerError(c)
	}

	err = db.CreateUser(creds.Email, hash, salt)
	if err != nil {
		log.WithErr(err).Error("Failed to create user account with email %s", creds.Email)
	}

	// TODO Send verification email

	return responder.Ok(c)
}

func completeUserProfile(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := newSession(c)

	idText := c.Params("id")
	id, _ := strconv.Atoi(idText)

	user := &models.User{}
	err := c.BodyParser(user)
	if err != nil {
		log.WithErr(err).Error("Failed to parse body")
		return responder.BadRequest(c, "Failed to parse user profile")
	}

	user.ID = uint(id)
	user.Verified = true

	user, err = db.UpdateUser(user)
	if err != nil {
		log.WithErr(err).Error("Failed to update user")
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, user)
}

func getUserById(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := newSession(c)

	idText := c.Params("id")
	id, _ := strconv.Atoi(idText)

	user, err := db.GetUserById(id)
	if err != nil {
		log.WithErr(err).Error("Failed to get user by ID %v", id)
	}

	if user == nil {
		return responder.NotFound(c)
	}

	return responder.OkWithData(c, user)
}
