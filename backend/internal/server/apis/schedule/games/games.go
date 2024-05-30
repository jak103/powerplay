package games

import (
	"encoding/csv"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/algorithms/round_robin"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"io"
	"mime/multipart"
	"strings"
)

var numberOfGamesPerTeam int

type Body struct {
	seasonID  uint
	algorithm string
	iceTimes  []string
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/games", auth.Authenticated, handleGenerate)
}

func handleGenerate(c *fiber.Ctx) error {
	numberOfGamesPerTeam = 10
	log.Info("Reading Body\n")

	body, err := readBody(c)
	if err != nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := body.seasonID
	algorithm := body.algorithm
	iceTimes := body.iceTimes

	// Read leagues from db
	logger := locals.Logger(c)
	session := db.GetSession(c)

	leagues, err := session.GetLeaguesBySeason(seasonID)
	if err != nil {
		logger.WithErr(err).Alert("Failed to get leagues for season %v the database", seasonID)
		return err
	}

	var games []structures.Game
	if algorithm == "round_robin" {
		games, err = round_robin.RoundRobin(leagues, iceTimes, numberOfGamesPerTeam)
	} else {
		return responder.BadRequest(c, fiber.StatusBadRequest, errors.New("invalid algorithm").Error())
	}
	// check for error after any of the algorithms is done
	if err != nil {
		return responder.InternalServerError(c, err)
	}

	// TODO save to db

	// TODO generate analysis

	return responder.Ok(c, games)
}

func readBody(c *fiber.Ctx) (Body, error) {

	type Dto struct {
		SeasonID  uint   `json:"season_id"`
		Algorithm string `json:"algorithm"`
	}

	var dto Dto
	if err := c.BodyParser(&dto); err != nil {
		return Body{}, err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return Body{}, err
	}
	iceTimes, err := getIceTimes(*file)
	if err != nil {
		return Body{}, err
	}

	body := Body{
		seasonID:  dto.SeasonID,
		algorithm: dto.Algorithm,
		iceTimes:  iceTimes,
	}

	return body, nil
}

func getIceTimes(file multipart.FileHeader) ([]string, error) {
	var iceTimes []string
	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		return iceTimes, err
	}
	defer func(uploadedFile multipart.File) {
		err := uploadedFile.Close()
		if err != nil {
			log.Error("Error closing file: %v", err)
		}
	}(uploadedFile)

	// Read the contents of the file
	csvContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		return iceTimes, err
	}
	reader := csv.NewReader(strings.NewReader(string(csvContent)))
	records, err := reader.ReadAll()
	if err != nil {
		return iceTimes, err
	}
	// get the headers
	headers := records[0]
	if headers[0] != "date" || headers[1] != "time" {
		return iceTimes, errors.New("invalid CSV file")
	}
	records = records[1:] // Skip the header
	for _, record := range records {
		iceTimes = append(iceTimes, record[0]+" "+record[1])
	}
	return iceTimes, nil
}
