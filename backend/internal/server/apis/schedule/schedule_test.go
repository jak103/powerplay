package schedule

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

// TODO mock the db.GetSession function

// Mocked types for testing
type MockSession struct {
	mock.Mock
}

type MockDB struct {
	mock.Mock
}

// Mock implementation for GetSession
func (m *MockDB) GetSession(c *fiber.Ctx) *MockSession {
	args := m.Called(c)
	return args.Get(0).(*MockSession)
}

// Mock implementation for GetGames
func (m *MockSession) GetGames(seasonID uint) (*[]Game, error) {
	args := m.Called(seasonID)
	return args.Get(0).(*[]Game), args.Error(1)
}

// Define a mock game struct for the purposes of testing
type Game struct {
	ID   int
	Name string
	// Define other relevant fields
}

// Mock implementation for SaveGames
func (m *MockSession) SaveGames(games []models.Game) (int, error) {
	args := m.Called(games)
	return args.Int(0), args.Error(1)
}

type CreateRequestBody struct {
	seasonID             uint
	iceTimes             int
	numberOfGamesPerTeam int
	// Add other fields as needed
}

func TestOptimizeGames(t *testing.T) {
	t.Run("Test successful optimization", func(t *testing.T) {
		// Mock the db module
		mockDB := new(MockDB)
		mockSession := new(MockSession)
		mockDB.On("GetSession", mock.Anything).Return(mockSession)

		// Mock the db.GetGames function
		seasonID := uint(123)
		mockGames := []Game{
			{ID: 1, Name: "Game 1"},
			{ID: 2, Name: "Game 2"},
		}
		mockSession.On("GetGames", seasonID).Return(&mockGames, nil)

		// Create Fiber app
		app := fiber.New()

		// Define handler function to call in Fiber app
		app.Put("/schedule", func(c *fiber.Ctx) error {
			return handleOptimizeGames(c)
		})

		// Create mock request body
		requestBody := `{"season_id": 123}`

		// Perform request using Fiber's Test function
		req := httptest.NewRequest("PUT", "/schedule", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Read the response body
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				t.Errorf("Error closing response body: %v", err)
			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		// Parse the response body
		var response struct {
			TeamStats []byte `json:"team_stats"`
			SeasonID  uint   `json:"season_id"`
		}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		// Assertions
		assert.NotNil(t, response.TeamStats)
		assert.Equal(t, seasonID, response.SeasonID)

		// Assert that expected functions were called
		mockDB.AssertExpectations(t)
		mockSession.AssertExpectations(t)
	})

	t.Run("Test failed optimization", func(t *testing.T) {
		t.Skip("Need to implement test")
	})
}

func TestCreateGames(t *testing.T) {

	t.Run("Test successful creation", func(t *testing.T) {
		// Mock setup
		mockDB := new(MockDB)
		mockSession := new(MockSession)
		mockDB.On("GetSession", mock.Anything).Return(mockSession)

		// Mock successful SaveGames call
		mockSession.On("SaveGames", mock.Anything).Return(1, nil)

		// Create Fiber app
		app := fiber.New()

		// Define handler function to call in Fiber app
		app.Post("/schedule", func(c *fiber.Ctx) error {
			return handleCreateGames(c)
		})

		// Prepare request body
		requestBody := `{"season_id": 123, "ice_times": 3, "number_of_games_per_team": 4}`

		// Perform request using Fiber's Test function
		req := httptest.NewRequest("POST", "/schedule", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse the response body
		var response struct {
			TeamStats []byte `json:"team_stats"`
			SeasonID  uint   `json:"season_id"`
		}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		// Assertions
		assert.NotNil(t, response.TeamStats)
		assert.Equal(t, uint(123), response.SeasonID)

		// Assert that expected functions were called
		mockDB.AssertExpectations(t)
		mockSession.AssertExpectations(t)
	})

	t.Run("Test failed creation", func(t *testing.T) {
		t.Skip("Need to implement test")
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
		assert.Equal(t, "pair-swap", body.optimizer)
		assert.Equal(t, 5, body.numberOfGamesPerTeam)
		assert.Empty(t, body.iceTimes)
	})

	t.Run("Test failed read body", func(t *testing.T) {
		t.Skip("Need to implement test")
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
		t.Skip("Need to implement test")
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
		t.Skip("Need to implement test")
	})
}

// Helper functions
func getCreateRequest(app *fiber.App) (*http.Request, error) {
	// Mock Fiber app
	app.Post("/schedule", handleCreateGames)

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
	_ = multipartWriter.WriteField("optimizer", "pair-swap")
	_ = multipartWriter.WriteField("number_of_games_per_team", "5")

	// Close the multipart writer
	if err := multipartWriter.Close(); err != nil {
		return nil, err
	}

	// Make a request to the endpoint with the multipart form data
	request := httptest.NewRequest("POST", "/schedule", &b)
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
