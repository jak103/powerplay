package parser

import (
	"fmt"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
	"os"

	"gopkg.in/yaml.v3"
)

func SeasonConfig(seasonName string) (*models.SeasonConfig, error) {
	configBytes, err := os.ReadFile(fmt.Sprintf("input/%s_config.yml", seasonName))
	if err != nil {
		configBytes, err = os.ReadFile(fmt.Sprintf("../input/%s_config.yml", seasonName))
		if err != nil {
			log.Error("Error reading file: %v\n", err)
			return nil, err
		}
	}

	log.Info("Read %v bytes\n", len(configBytes))

	seasonConfig := &models.SeasonConfig{}
	err = yaml.Unmarshal(configBytes, &seasonConfig)
	if err != nil {
		log.Error("Error parsing yaml: %v\n", err)
		return nil, err
	}

	return seasonConfig, nil
}
