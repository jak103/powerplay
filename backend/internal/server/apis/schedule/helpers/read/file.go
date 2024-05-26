package read

import (
	"fmt"
	"github.com/jak103/powerplay/internal/utils/log"
	"gopkg.in/yaml.v3"
	"os"
)

func IceTimes(seasonName string) ([]string, error) {
	fileBytes, err := os.ReadFile(fmt.Sprintf("uploads/%s.yml", seasonName))
	if err != nil {
		fileBytes, err = os.ReadFile(fmt.Sprintf("../uploads/%s.yml", seasonName))
		if err != nil {
			log.Error("Error reading file: %v\n", err)
			return nil, err
		}
	}

	log.Info("Read %v bytes\n", len(fileBytes))

	var times struct {
		IceTimes []string `yaml:"ice_time"`
	}

	err = yaml.Unmarshal(fileBytes, &times)

	if err != nil {
		log.Error("Error parsing yaml: %v\n", err)
		return nil, err
	}

	return times.IceTimes, nil
}
