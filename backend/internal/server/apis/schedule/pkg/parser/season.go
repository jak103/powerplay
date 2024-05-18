package parser

import (
	"fmt"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"os"

	"gopkg.in/yaml.v3"
)

func SeasonConfig(season string) (*models.SeasonConfig, error) {
	configBytes, err := os.ReadFile(fmt.Sprintf("input/%s_config.yml", season))
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil, err
	}

	fmt.Printf("Read %v bytes\n", len(configBytes))

	seasonConfig := &models.SeasonConfig{}
	err = yaml.Unmarshal(configBytes, &seasonConfig)
	if err != nil {
		fmt.Println("Error parsing yaml", err)
		return nil, err
	}

	return seasonConfig, nil
}
