package csv

import (
	"os"

    "github.com/jak103/powerplay/internal/utils/log"
	"github.com/gocarina/gocsv"
)

func GenerateCsv[T any](games []T, filename string) {
	log.Info("Writing CSV: %s", filename)

	csvGames, err := gocsv.MarshalBytes(games)
	if err != nil {
		log.Error("Failed to marshal games %v", err)
	}

	log.Info("Data marshaled, now writing")

	err = os.WriteFile(filename, csvGames, 0644)
	if err != nil {
		log.Error("Failed to write games %v", err)
	}
	log.Info("Done writing")
}
