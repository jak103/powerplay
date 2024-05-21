package schedule

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Schedule struct {
	Date     time.Time
	Time     time.Time
	Duration time.Duration
}

func readCSV() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a CSV file as the first argument")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening CSV file: ", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing CSV file: ", err)
			return
		}
	}(file)

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file: ", err)
		return
	}

	var dateTimes []Schedule

	for _, line := range lines {
		dateStr := line[0]
		timeStr := line[1]

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Error parsing date: ", err)
			continue
		}

		timeOfDay, err := time.Parse("15:04:05", timeStr)
		if err != nil {
			fmt.Println("Error parsing time: ", err)
			continue
		}

		// Create Schedule struct
		duration := 75 * time.Minute
		schedule := Schedule{
			Date:     date,
			Time:     timeOfDay,
			Duration: duration,
		}

		dateTimes = append(dateTimes, schedule)
	}
}
