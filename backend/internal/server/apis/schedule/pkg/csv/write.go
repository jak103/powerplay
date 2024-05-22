package csv

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jak103/powerplay/internal/utils/log"
)

func GenerateCsv[T any](games []T, filename string) error {
	path := "output/" + filename
	log.Info("Writing CSV: %s", path)

	csvGames, err := gocsv.MarshalBytes(games)
	if err != nil {
		log.Error("Failed to marshal games %v", err)
		return err
	}

	log.Info("Data marshaled, now writing")

	err = os.WriteFile(path, csvGames, 0644)
	if err != nil {
		path = "../output/" + filename
		err = os.WriteFile(path, csvGames, 0644)
		if err != nil {
			log.Error("Failed to write games %v", err)
			return err
		}
	}
	log.Info("Done writing")
	return nil
}
