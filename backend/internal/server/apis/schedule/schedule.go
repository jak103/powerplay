package schedule

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/round_robin"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/optimize"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

// The following endpoints are for the schedule creation
// - first the caller will call handleCreateGames to create the games
// - then the caller can call handleOptimizeGames to optimize the schedule as many times as they want
// - then the caller will call handleSaveGames to save the games to the database

type Body struct {
	seasonID             uint
	venueID              uint
	optimizer            string
	iceTimes             []string
	numberOfGamesPerTeam int
}

type response struct {
	TeamStats []structures.TeamStats
	SeasonID  uint `json:"seasonId"`
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule", auth.Authenticated, handleCreateGames)
	apis.RegisterHandler(fiber.MethodPut, "/schedule", auth.Authenticated, handleOptimizeGames)
}

func handleOptimizeGames(c *fiber.Ctx) error {
	type Dto struct {
		SeasonID uint `json:"season_id"`
	}
	var dto Dto
	body, err := readBody(c)
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

	optimizer := body.optimizer

	if optimizer == "pair-swap" {
		optimize.PairOptimizeSchedule(*games)
	} else if optimizer == "set-swap" {
		optimize.SetOptimizeSchedule(*games)
	} else {
		return responder.BadRequest(c, fiber.StatusBadRequest, fmt.Sprintf("Invalid optimizer %v", optimizer))
	}

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
	venueID := body.venueID
	optimizer := body.optimizer
	iceTimes := body.iceTimes
	numberOfGamesPerTeam := body.numberOfGamesPerTeam

	// Read leagues from db
	logger := locals.Logger(c)
	session := db.GetSession(c)

	leagues, err := session.GetLeaguesBySeason(int(seasonID))
	if err != nil {
		logger.WithErr(err).Error("Failed to get leagues for season %v the database", seasonID)
		return responder.InternalServerError(c, err)
	}

	if leagues == nil {
		return responder.BadRequest(c, errors.New("no league for the season").Error())
	}

	venue, err := session.GetVenueById(venueID)
	if err != nil {
		logger.WithErr(err).Error("Failed to get venue with id of %v the database", venueID)
		return responder.InternalServerError(c, err)
	}

	if venue == nil {
		return responder.BadRequest(c, errors.New("no venue for the season").Error())
	}

	var games []models.Game
	games, err = round_robin.RoundRobin(leagues, iceTimes, numberOfGamesPerTeam, *venue)
	// check for error after any of the algorithms is done
	if err != nil {
		return responder.InternalServerError(c, err.Error())
	}

	assignLockerRooms(games)

	// optimize the schedule 10 times and pick the one with the best score
	seasonStats, ts := analysis.RunTimeAnalysis(games)
	bestScore := seasonStats.Score
	bestGames := make([]models.Game, len(games))
	for i, game := range games {
		bestGames[i] = game
	}
	for i := 0; i < 10; i++ {
		if optimizer == "pair-swap" {
			optimize.PairOptimizeSchedule(games)
		} else if optimizer == "set-swap" {
			optimize.SetOptimizeSchedule(games)
		} else {
			return responder.BadRequest(c, fiber.StatusBadRequest, errors.New(fmt.Sprintf("Invalid optimizer %v", optimizer)).Error())
		}

		seasonStats, ts = analysis.RunTimeAnalysis(games)

		if seasonStats.Score > bestScore {
			for i, game := range games {
				bestGames[i] = game
			}
			bestScore = seasonStats.Score
		}
	}

	seasonStats, ts = analysis.RunTimeAnalysis(bestGames)

	// save to db
	_, err = session.SaveGames(bestGames)
	if err != nil {
		log.Info("Failed to save games to the database\n")
		return responder.InternalServerError(c, err.Error())
	}

	data := response{
		TeamStats: analysis.Serialize(ts),
		SeasonID:  seasonID,
	}

	return responder.OkWithData(c, data)
}

func readBody(c *fiber.Ctx) (Body, error) {

	// keys
	// - season_id
	// - venue_id
	// - algorithm
	// - file

	dto := struct {
		SeasonID             uint   `json:"season_id"`
		VenueID              uint   `json:"venue_id"`
		Optimizer            string `json:"optimizer"`
		NumberOfGamesPerTeam int    `json:"number_of_games_per_team"`
	}{}

	if err := c.BodyParser(&dto); err != nil {
		log.Error("Error reading the body")
		return Body{}, responder.BadRequest(c, "Failed to parse body of request")
	}

	file, err := c.FormFile("ice_times")
	if err != nil {
		log.Error("Error reading the file")
		return Body{}, err
	}

	iceTimes, err := getIceTimes(*file)
	if err != nil {
		log.Error("Error reading the ice times")
		return Body{}, err
	}

	body := Body{
		seasonID:             dto.SeasonID,
		venueID:              dto.VenueID,
		optimizer:            dto.Optimizer,
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
		if analysis.IsEarlyGame(game.Start.Hour(), game.Start.Minute()) {
			games[i].HomeTeamLockerRoom = "3"
			games[i].AwayTeamLockerRoom = "1"
		} else {
			games[i].HomeTeamLockerRoom = "5"
			games[i].AwayTeamLockerRoom = "2"
		}
	}
}
