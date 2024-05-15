package csv

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func GenerateCsv[T any](games []T, filename string) {
	fmt.Println("Writing CSV:", filename)

	csvGames, err := gocsv.MarshalBytes(games)
	if err != nil {
		fmt.Printf("Failed to marshal games %v", err)
	}

	fmt.Println("Data marshaled, now writing")

	err = os.WriteFile(filename, csvGames, 0644)
	if err != nil {
		fmt.Printf("Failed to write games %v", err)
	}
	fmt.Println("Done writing")
}
