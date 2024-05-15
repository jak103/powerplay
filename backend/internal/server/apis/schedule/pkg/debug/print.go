package debug

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func Print(in any) {
	leaguesText, _ := yaml.Marshal(in)
	fmt.Printf("%v\n", string(leaguesText))
}
