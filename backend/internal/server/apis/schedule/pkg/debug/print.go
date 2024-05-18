package debug

import (
	"github.com/jak103/powerplay/internal/utils/log"
	"gopkg.in/yaml.v3"
)

func Print(in any) {
	leaguesText, _ := yaml.Marshal(in)
	log.Info("%v\n", string(leaguesText))
}
