package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type createResponse struct {
	Email  string `json:"email"`
	UserId int    `json:"user_id"`
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/user", auth.Authenticated, getCurrentUser)
	apis.RegisterHandler(fiber.MethodPost, "/user", auth.Public, createUserAccount)

}

func removeFormat(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
}

func validateUser(u *models.User) error {

	//Check data field has been filled for all values
	values := reflect.ValueOf(*u)
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).String()
		if v == "" {
			return errors.New("data field is empty")
		}
	}

	//Validate email has an @ in middle
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("email is invalid")
	}
	//Validate phone number is 10 digit int
	u.Phone = removeFormat(u.Phone)
	if _, err := strconv.Atoi(u.Phone); err != nil || len(u.Phone) != 10 {
		return errors.New("phone number is invalid")
	}
	//Validate skill level is an at least 0
	if u.SkillLevel < 0 {
		return errors.New("skill level is negative")
	}
	return nil
}

func getCurrentUser(c *fiber.Ctx) error {
	return nil
}

func createUserAccount(c *fiber.Ctx) error {

	// verify the request
	log := locals.Logger(c)

	// anonymous struct to hold the request data
	creds := struct {
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		Phone       string    `json:"phone"`
		SkillLevel  int       `json:"skill_level"`
		DateOfBirth time.Time `json:"date_of_birth"`
	}{}

	err := c.BodyParser(&creds)
	if err != nil {
		log.WithErr(err).Error("Failed to parse user creation request")
		return responder.BadRequest(c, "Failed to parse user creation request")
	}

	// marshal the request data into a User struct
	user := models.User{
		FirstName:   creds.FirstName,
		LastName:    creds.LastName,
		Email:       creds.Email,
		Password:    creds.Password,
		Phone:       creds.Phone,
		SkillLevel:  creds.SkillLevel,
		DateOfBirth: creds.DateOfBirth,
	}

	// validate the request
	err = validateUser(&user)
	if err != nil {
		log.WithErr(err).Error(err.Error())
		return responder.BadRequest(c, err.Error())
	}

	// parse actual user object
	u := &models.User{}
	err = c.BodyParser(&u)
	if err != nil {
		log.WithErr(err).Error(err.Error())
		return responder.BadRequest(c, err.Error())
	}

	u.Password = hashPassword(creds.Password)

	// write to database
	db := db.GetSession(c)
	log.Debug("Creating user %s", user.Email)
	u, result := db.CreateUser(u)
	if result != nil {
		log.WithErr(err).Error(result.Error())
		return responder.InternalServerError(c, result.Error())
	}

	// response
	createdUserResponse := createResponse{
		Email:  user.Email,
		UserId: int(u.ID),
	}

	return responder.OkWithData(c, createdUserResponse)

}

func hashPassword(password string) string {
	key := config.Vars.PasswordKey
	keyBytes := []byte(key)

	passwordBytes := []byte(password)
	mac := hmac.New(sha256.New, keyBytes)
	mac.Write(passwordBytes)
	hashedPassword := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashedPassword)
}
