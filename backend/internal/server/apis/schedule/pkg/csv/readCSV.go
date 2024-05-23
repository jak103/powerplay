package csv

import (
	"encoding/csv"
	"github.com/jak103/powerplay/internal/utils/log"
	"os"
	"time"
)

type Schedule struct {
	Date     time.Time
	Time     time.Time
	Duration time.Duration
}

//	 Reads CSV with the following format:
//		[Date] [Time]
//		YYYY-MM-DD, HH:MM:SS
//		Returns an array of Schedule structs
func readCSV(filename string) ([]Schedule, error) {
	const GameDuration = 75
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing CSV file: %v", err)
			return
		}
	}(file)

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dateTimes []Schedule

	log.Info("Reading CSV file into Schedule Struct")
	for _, line := range lines {
		dateStr := line[0]
		timeStr := line[1]

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Error("Error parsing date: %v", err)
			continue
		}

		timeOfDay, err := time.Parse("15:04:05", timeStr)
		if err != nil {
			log.Error("Error parsing time: %v", err)
			continue
		}

		// Create Schedule struct
		duration := GameDuration * time.Minute
		schedule := Schedule{
			Date:     date,
			Time:     timeOfDay,
			Duration: duration,
		}

		dateTimes = append(dateTimes, schedule)
	}
	return dateTimes, nil
}
