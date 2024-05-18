package parser

import (
	"fmt"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"os"
    "github.com/jak103/powerplay/internal/utils/log"

	"gopkg.in/yaml.v3"
)

func SeasonConfig(season string) (*models.SeasonConfig, error) {
	configBytes, err := os.ReadFile(fmt.Sprintf("input/%s_config.yml", season))
	if err != nil {
        log.Error("Error reading file: %v\n", err)
		return nil, err
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
