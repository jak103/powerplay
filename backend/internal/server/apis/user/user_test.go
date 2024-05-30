package user

import (
	"testing"

	"github.com/jak103/powerplay/internal/models"
)

func TestValidateUser(t *testing.T) {
	var tests = []struct {
		name  string
		input models.User
		want  string
	}{
		{"User should be valid", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", Password: "password", Phone: "123-456-7890", SkillLevel: 3, Role: nil}, ""},
		{"User should be valid", models.User{FirstName: "bill", LastName: "bob", Email: "bb@yahoo.com", Password: "password123@bob", Phone: "(123) 456-7890", SkillLevel: 39, Role: nil}, ""},
		{"First name should be empty", models.User{FirstName: "", LastName: "Doe", Email: "doe@gmail.com", Password: "password", Phone: "123-456-7890", SkillLevel: 3, Role: nil}, "data field is empty"},
		{"Password should be empty", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", Password: "", Phone: "123-456-7890", SkillLevel: 3, Role: nil}, "data field is empty"},
		{"Email should be invalid", models.User{FirstName: "John", LastName: "Doe", Email: "@gmail.com", Password: "password", Phone: "123-456-7890", SkillLevel: 3, Role: nil}, "email is invalid"},
		{"Email should be invalid", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com@", Password: "password", Phone: "123-456-7890", SkillLevel: 3, Role: nil}, "email is invalid"},
		{"Phone number should be invalid", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", Password: "password", Phone: "123-456-7890a", SkillLevel: 3, Role: nil}, "phone number is invalid"},
		{"Phone number should be invalid", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", Password: "password", Phone: "123-456-78905", SkillLevel: 3, Role: nil}, "phone number is invalid"},
		{"Skill level should be invalid", models.User{FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", Password: "password", Phone: "123-456-7890", SkillLevel: -4, Role: nil}, "skill level is negative"},
	}
	// The execution loop
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			err := validateUser(&tt.input)
			println("validateUser(&tt.input): ", err)
			if (err == nil && tt.want != "") || (err != nil && err.Error() != tt.want) {
				t.Errorf("validateUser(%v) = %v, want %v", tt.input, err, tt.want)
			}
		})
	}
}
