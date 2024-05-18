package debug

import (
	"gopkg.in/yaml.v3"
    "github.com/jak103/powerplay/internal/utils/log"
)

func Print(in any) {
	leaguesText, _ := yaml.Marshal(in)
	log.Info("%v\n", string(leaguesText))
}
