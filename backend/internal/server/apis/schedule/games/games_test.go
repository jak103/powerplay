package games

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
	"time"
)

// TODO add dockertest to be able to test the functions that interact with the database

func TestOptimizeGames(t *testing.T) {
	t.Run("Test successful optimization", func(t *testing.T) {
		// Mock Fiber app
		app := fiber.New()
		app.Put("/schedule/auto/optimize", handleOptimizeGames)

		// Prepare request body
		requestBody := `{"season_id": 123}`

		// Make a request to the endpoint
		request := httptest.NewRequest("PUT", "/schedule/auto/optimize", strings.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(request)

		// Check the response
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse the response body
		var response struct {
			TeamStats []byte `json:"team_stats"`
			SeasonID  uint   `json:"season_id"`
		}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		// You can add more specific assertions based on the expected response data
		assert.NotNil(t, response.TeamStats)
		assert.Equal(t, uint(123), response.SeasonID)
	})

	t.Run("Test failed optimization", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestCreateGames(t *testing.T) {

	t.Run("Test successful creation", func(t *testing.T) {
		app := fiber.New()
		request, err := getCreateRequest(app)
		if err != nil {
			t.Errorf("Failed to create request: %v", err)
			return
		}
		resp, err := app.Test(request)

		// Check the response
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
			return
		}
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Test failed creation", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestReadBody(t *testing.T) {

	t.Run("Test successful read body", func(t *testing.T) {
		app := fiber.New()
		request, err := getCreateRequest(app)

		fastRequest := convertHTTPToFastHTTP(request)

		// Simulate request using Fiber's context acquisition method
		ctx := app.AcquireCtx(fastRequest)
		defer app.ReleaseCtx(ctx)

		// Call the readBody function with the created context
		body, err := readBody(ctx)

		// Check the result
		assert.NoError(t, err)
		assert.NotNil(t, body)
		assert.Equal(t, uint(123), body.seasonID)
		assert.Equal(t, "round_robin", body.algorithm)
		assert.Equal(t, 5, body.numberOfGamesPerTeam)
		assert.Empty(t, body.iceTimes)
	})

	t.Run("Test failed read body", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestGetIceTimes(t *testing.T) {

	t.Run("Test successful get ice times", func(t *testing.T) {
		// Create a mock multipart file
		fileHeader, _, err := getMultiPartFile()
		if err != nil {
			t.Fatalf("Error creating mock multipart file: %v", err)
		}

		// Call the getIceTimes function
		iceTimes, err := getIceTimes(*fileHeader)

		// Check the result
		assert.Error(t, err)
		assert.Nil(t, iceTimes)
	})

	t.Run("Test failed get ice times", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestAssignLockerRooms(t *testing.T) {

	t.Run("Test successful assign", func(t *testing.T) {
		// Mock games
		games := getGames()

		// Call assignLockerRooms function
		assignLockerRooms(games)

		// Check the result
		for _, game := range games {
			// Assert that locker rooms are not empty
			assert.NotEmpty(t, game.HomeTeamLockerRoom)
			assert.NotEmpty(t, game.AwayTeamLockerRoom)

			// Determine expected locker rooms based on the time of the game
			expectedHomeLockerRoom := ""
			expectedAwayLockerRoom := ""
			if isEarlyGame(game.Start.Hour(), game.Start.Minute()) {
				expectedHomeLockerRoom = "3"
				expectedAwayLockerRoom = "1"
			} else {
				expectedHomeLockerRoom = "5"
				expectedAwayLockerRoom = "2"
			}

			// Assert that the assigned locker rooms match the expected values
			assert.Equal(t, expectedHomeLockerRoom, game.HomeTeamLockerRoom)
			assert.Equal(t, expectedAwayLockerRoom, game.AwayTeamLockerRoom)
		}
	})

	t.Run("Test failed assign", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

// Helper functions
func getCreateRequest(app *fiber.App) (*http.Request, error) {
	// Mock Fiber app
	app.Post("/schedule/auto/games", handleCreateGames)

	// Get the multipart file
	fileHeader, body, err := getMultiPartFile()
	if err != nil {
		return nil, err
	}

	// Create a new multipart form writer
	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	// Write the file part
	fileWriter, err := multipartWriter.CreateFormFile("ice_times", fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(fileWriter, body); err != nil {
		return nil, err
	}

	// Write other fields
	_ = multipartWriter.WriteField("seasonID", "123")
	_ = multipartWriter.WriteField("algorithm", "round_robin")
	_ = multipartWriter.WriteField("number_of_games_per_team", "5")

	// Close the multipart writer
	if err := multipartWriter.Close(); err != nil {
		return nil, err
	}

	// Make a request to the endpoint with the multipart form data
	request := httptest.NewRequest("POST", "/schedule/auto/games", &b)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	return request, nil
}

func convertHTTPToFastHTTP(req *http.Request) *fasthttp.RequestCtx {
	// Create a new fasthttp request
	fastReq := fasthttp.AcquireRequest()

	// Convert the http request to fasthttp request
	fastReq.Header.SetMethodBytes([]byte(req.Method))
	fastReq.SetRequestURIBytes([]byte(req.URL.RequestURI()))
	fastReq.Header.SetContentType(req.Header.Get("Content-Type"))

	// Set the body if the content length is greater than 0
	if req.ContentLength > 0 {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			// Handle error
		}
		fastReq.SetBody(body)
	}

	// Create a new fasthttp request context
	ctx := &fasthttp.RequestCtx{}

	// Set the request to the context
	ctx.Init(fastReq, nil, nil)

	// Return the fasthttp request context
	return ctx
}

func isEarlyGame(hour, minute int) bool {
	if hour < 20 {
		return true
	}
	switch hour {
	case 20:
		return true
	case 21:
		return minute <= 15
	case 22, 23:
		return false
	}
	return false
}

func getGames() []models.Game {
	// Define start times for each game
	startTimes := []time.Time{
		time.Date(2024, time.June, 11, 18, 0, 0, 0, time.UTC),  // Game 1 starts at 6:00 PM
		time.Date(2024, time.June, 11, 19, 30, 0, 0, time.UTC), // Game 2 starts at 7:30 PM
		time.Date(2024, time.June, 11, 21, 0, 0, 0, time.UTC),  // Game 3 starts at 9:00 PM
		time.Date(2024, time.June, 11, 22, 30, 0, 0, time.UTC), // Game 4 starts at 10:30 PM
		time.Date(2024, time.June, 11, 23, 59, 0, 0, time.UTC), // Game 5 starts at 11:59 PM
	}

	// Create games with corresponding start times
	var games []models.Game
	for i := 0; i < len(startTimes); i++ {
		games = append(games, models.Game{
			Start:      startTimes[i],
			HomeTeamID: uint(i*2 + 1),
			AwayTeamID: uint(i*2 + 2),
		})
	}
	return games
}

func getMultiPartFile() (*multipart.FileHeader, *bytes.Buffer, error) {
	// Data to be written into the CSV file
	csvData := `date,time
    5/20/24,20:30
    5/20/24,22:00
    5/21/24,20:45
    5/21/24,22:15
    5/22/24,20:45
    5/22/24,22:15
    5/23/24,20:45
    5/23/24,22:15
    5/24/24,21:45
    5/25/24,20:45
    5/25/24,22:15
    5/27/24,20:30
    5/27/24,22:00
    5/28/24,20:45`

	// Create a buffer to write the CSV data
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write the CSV data to the buffer
	err := writer.Write([]string{"date", "time"})
	if err != nil {
		return nil, nil, err
	}
	writer.Flush()

	// Write CSV data line by line
	csvLines := csv.NewReader(bytes.NewReader([]byte(csvData)))
	for {
		record, err := csvLines.Read()
		if err != nil {
			break
		}
		err = writer.Write(record)
		if err != nil {
			return nil, nil, err
		}
		writer.Flush()
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	// Create a part for the CSV file
	partWriter, err := multipartWriter.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="csv-file"; filename="data.csv"`},
		"Content-Type":        []string{"text/csv"},
	})
	if err != nil {
		return nil, nil, err
	}

	// Write the CSV data to the part
	if _, err := partWriter.Write(buf.Bytes()); err != nil {
		return nil, nil, err
	}

	// Close the multipart writer
	if err := multipartWriter.Close(); err != nil {
		return nil, nil, err
	}

	// Create a new multipart file header
	fileHeader := &multipart.FileHeader{
		Filename: "data.csv",
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"text/csv"},
		},
		// Put the content of the buffer into a file
		Size: int64(body.Len()),
	}
	return fileHeader, body, nil
}
