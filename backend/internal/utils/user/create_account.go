package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
)

type User struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name" `
	Email      string `json:"email"    `
	Password   string `json:"-"` // Password should never leave the backend
	Phone      string `json:"phone"    `
	SkillLevel int    `json:"skill_level"`
	UserName   string `json:"user_name"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Parse the JSON request and populate the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//Remove non-Alphanumerics like () and - from phone number
	user.Phone = removeFormat(user.Phone)
	//Check if data is valid
	err = validateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err), http.StatusBadRequest)
		return
	}

	// Validation passed, we can now create the account here

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func removeFormat(str string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
}

func validateUser(u User) error {
	values := reflect.ValueOf(u)
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).String()
		//Check data field has been filled for all values
		if v == "" {
			return fmt.Errorf("data field is empty")
		}
	}
	//Validate email has an @ in middle
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("%s is an invalid Email", u.Email)
	}
	//Validate phone number is 10 digit int
	if _, err := strconv.Atoi(u.Phone); err != nil || len(u.Phone) != 10 {
		return fmt.Errorf("%s is an invalid phone number", u.Phone)
	}
	//Validate skill level is an at least 0
	if u.SkillLevel < 0 {
		return fmt.Errorf("skill level must be positive, %d is negative", u.SkillLevel)
	}
	return nil
}
