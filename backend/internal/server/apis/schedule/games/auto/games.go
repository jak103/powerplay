package auto

import (
	"encoding/csv"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/algorithms/round_robin"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"io"
	"mime/multipart"
	"strings"
)

// The following endpoints are for the schedule creation
// - first the caller will call handleCreateGames to create the games
// - then the caller can call handleOptimizeGames to optimize the schedule as many times as they want
// - then the caller will call handleSaveGames to save the games to the database

type Body struct {
	seasonID             uint
	algorithm            string
	iceTimes             []string
	numberOfGamesPerTeam int
}

type response struct {
	TeamStats []structures.TeamStats
    SeasonID  uint                       `json:"seasonId"`
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/auto/games", auth.Authenticated, handleCreateGames)
	apis.RegisterHandler(fiber.MethodPut, "/schedule/auto/optimize", auth.Authenticated, handleOptimizeGames)
}

func handleOptimizeGames(c *fiber.Ctx) error {
	type Dto struct {
		SeasonID uint `json:"season_id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := dto.SeasonID
	// read from db
	session := db.GetSession(c)
	games, err := session.GetGames(seasonID)
	if err != nil {
		log.Info("Failed to get games from the database\n")
		return responder.InternalServerError(c, err)
	}

	round_robin.OptimizeSchedule(*games)
	// write to the db
	assignLockerRooms(*games)
	_, err = session.SaveGames(*games)
	if err != nil {
		log.Info("Failed to save games to the database\n")
		return responder.InternalServerError(c, err)
	}

	_, ts := analysis.RunTimeAnalysis(*games)
	data := response{
		TeamStats: analysis.Serialize(ts),
		SeasonID:  seasonID,
	}

	return responder.OkWithData(c, data)
}

func handleCreateGames(c *fiber.Ctx) error {
	log.Info("Reading Body\n")

	body, err := readBody(c)
	if err != nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := body.seasonID
	algorithm := body.algorithm
	iceTimes := body.iceTimes
	numberOfGamesPerTeam := body.numberOfGamesPerTeam

	// Read leagues from db
	logger := locals.Logger(c)
	session := db.GetSession(c)

	leagues, err := session.GetLeaguesBySeason(seasonID)
	if err != nil {
		logger.WithErr(err).Error("Failed to get leagues for season %v the database", seasonID)
		return responder.InternalServerError(c, err)
	}

	if leagues == nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, errors.New("no league for the season").Error())
	}

	var games []models.Game
	if algorithm == "round_robin" {
		games, err = round_robin.RoundRobin(leagues, iceTimes, numberOfGamesPerTeam)
	} else {
		return responder.BadRequest(c, fiber.StatusBadRequest, errors.New("invalid algorithm").Error())
	}
	// check for error after any of the algorithms is done
	if err != nil {
		return responder.InternalServerError(c, err)
	}

	assignLockerRooms(games)

	// save to db
	_, err = session.SaveGames(games)
	if err != nil {
		log.Info("Failed to save games to the database\n")
		return responder.InternalServerError(c, err)
	}

	_, ts := analysis.RunTimeAnalysis(games)

	data := response{
		TeamStats: analysis.Serialize(ts),
		SeasonID:  seasonID,
	}

	return responder.OkWithData(c, data)
}

func readBody(c *fiber.Ctx) (Body, error) {

	// keys
	// - season_id
	// - algorithm
	// - file

	type Dto struct {
		SeasonID             uint   `json:"season_id"`
		Algorithm            string `json:"algorithm"`
		NumberOfGamesPerTeam int    `json:"number_of_games_per_team"`
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
		seasonID:             dto.SeasonID,
		algorithm:            dto.Algorithm,
		iceTimes:             iceTimes,
		numberOfGamesPerTeam: dto.NumberOfGamesPerTeam,
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

func assignLockerRooms(games []models.Game) {
	// The algorithm is pretty simple.
	//For the early game, home team is locker room 3, and away is locker room 1.
	//For the late game home team is locker room 5, and away team is locker room 2.

	for i, game := range games {
		if round_robin.IsEarlyGame(game.Start.Hour(), game.Start.Minute()) {
			games[i].HomeTeamLockerRoom = "3"
			games[i].AwayTeamLockerRoom = "1"
		} else {
			games[i].HomeTeamLockerRoom = "5"
			games[i].AwayTeamLockerRoom = "2"
		}
	}
}
