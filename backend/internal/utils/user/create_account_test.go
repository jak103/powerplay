package user

import (
	"testing"
)

func TestValidateUser(t *testing.T) {
	var tests = []struct {
		name  string
		input User
		want  string
	}{
		{"User should be valid", User{"John", "Doe", "doe@gmail.com", "password", "123-456-7890", 3, "JD123"}, ""},
		{"User should be valid", User{"bill", "bob", "bb@yahoo.com", "password123@bob", "(123) 456-7890", 39, "JD1-23"}, ""},
		{"First name should be empty", User{"", "Doe", "doe@gmail.com", "password", "123-456-7890", 3, "JD123"}, "data field is empty"},
		{"Password should be empty", User{"John", "Doe", "doe@gmail.com", "", "123-456-7890", 3, "JD123"}, "data field is empty"},
		{"Email should be invalid", User{"John", "Doe", "@gmail.com", "password", "123-456-7890", 3, "JD123"}, "email is invalid"},
		{"Email should be invalid", User{"John", "Doe", "doe@gmail.com@", "password", "123-456-7890", 3, "JD123"}, "email is invalid"},
		{"Phone number should be invalid", User{"John", "Doe", "doe@gmail.com", "password", "123-456-7890a", 3, "JD123"}, "phone number is invalid"},
		{"Phone number should be invalid", User{"John", "Doe", "doe@gmail.com", "password", "123-456-78905", 3, "JD123"}, "phone number is invalid"},
		{"Skill level should be invalid", User{"John", "Doe", "doe@gmail.com", "password", "123-456-7890", -4, "JD123"}, "skill level is negative"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := validateUser(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
