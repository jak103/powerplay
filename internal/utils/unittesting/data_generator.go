package unittesting

import (
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/utils/formatters"
)

func GenerateUniqueName(existingNames map[string]bool) string {
	var name string
	for {
		name = formatters.CapitalizeFirstLetter(faker.Word())
		if !existingNames[name] {
			existingNames[name] = true
			break
		}
	}
	return name
}
