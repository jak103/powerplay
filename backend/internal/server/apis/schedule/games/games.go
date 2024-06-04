package games

import (
	"encoding/csv"
	"errors"
	"github.com/jak103/powerplay/internal/models"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/algorithms/round_robin"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

var numberOfGamesPerTeam int

type Body struct {
	seasonID  uint
	algorithm string
	iceTimes  []string
}

type response struct {
	TeamStats []structures.TeamStats
}

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/games", auth.Authenticated, handleGenerate)
	apis.RegisterHandler(fiber.MethodPost, "/schedule/save", auth.Authenticated, handleSave)
	apis.RegisterHandler(fiber.MethodPut, "/schedule/optimize", auth.Authenticated, handleOptimize)
}

func handleOptimize(c *fiber.Ctx) error {
	type Dto struct {
		Games []structures.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games
	round_robin.OptimizeSchedule(games)
	_, ts := analysis.RunTimeAnalysis(games)

	data := response{
		TeamStats: analysis.Serialize(ts),
	}

	return responder.OkWithData(c, data)
}

func handleSave(c *fiber.Ctx) error {
	type Dto struct {
		Games []structures.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games

	session := db.GetSession(c)
	dbGames := mapGameStructToGameModel(games)
	_, err = session.SaveGames(dbGames)
	if err != nil {
		log.Error("Failed to save games to the database")
		return responder.InternalServerError(c, err)

	}
	return responder.Ok(c)
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
		logger.WithErr(err).Error("Failed to get leagues for season %v the database", seasonID)
		return responder.InternalServerError(c, err)
	}

	if leagues == nil {
		return responder.BadRequest(c, fiber.StatusBadRequest, errors.New("no league for the season").Error())
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

	assignLockerRooms(games)

	_, ts := analysis.RunTimeAnalysis(games)

	data := response{
		TeamStats: analysis.Serialize(ts),
	}

	return responder.OkWithData(c, data)
}

func readBody(c *fiber.Ctx) (Body, error) {

	// keys
	// season_id
	// algorithm
	// file

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

func assignLockerRooms(games []structures.Game) {
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

func mapGameStructToGameModel(games []structures.Game) []models.Game {
	var gameModels []models.Game
	// TODO deal with the VenueID
	// TODO deal with the correlationId
	// TODO deal with roster
	for _, game := range games {
		gameModels = append(gameModels, models.Game{
			DbModel:  game.DbModel,
			SeasonID: game.SeasonID,
			Start:    game.Start,
			Venue: models.Venue{
				Name:        "George S. Eccles Ice Center",
				Address:     "2825 N 200 E North Logan, UT 84341 United States",
				LockerRooms: []string{"1", "2", "3", "4", "5"},
			},
			VenueID:             0,
			Status:              models.SCHEDULED,
			HomeTeam:            game.HomeTeam,
			HomeTeamID:          game.HomeTeamID,
			HomeTeamRoster:      game.HomeTeamRoster,
			HomeTeamRosterID:    game.HomeTeamRosterID,
			HomeTeamLockerRoom:  game.HomeTeamLockerRoom,
			HomeTeamShotsOnGoal: game.HomeTeamShotsOnGoal,
			HomeTeamScore:       game.HomeTeamScore,
			AwayTeam:            game.AwayTeam,
			AwayTeamID:          game.AwayTeamID,
			AwayTeamRoster:      game.AwayTeamRoster,
			AwayTeamRosterID:    game.AwayTeamRosterID,
			AwayTeamLockerRoom:  game.AwayTeamLockerRoom,
			AwayTeamShotsOnGoal: game.AwayTeamShotsOnGoal,
			AwayTeamScore:       game.AwayTeamScore,
			ScoreKeeper:         game.ScoreKeeper,
			ScoreKeeperID:       game.ScoreKeeperID,
			PrimaryReferee:      game.PrimaryReferee,
			PrimaryRefereeID:    game.PrimaryRefereeID,
			SecondaryReferee:    game.SecondaryReferee,
			SecondaryRefereeID:  game.SecondaryRefereeID,
		})
	}
	return gameModels
}
