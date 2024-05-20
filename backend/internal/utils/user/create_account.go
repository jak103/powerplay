package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
	/*I removed non-alphanumerics and space from the number so that
	  we can store it without any formatting*/
	user.Phone = removeFormat(user.Phone)
	//Check if data is valid
	message := validateUser(user)
	if message != "" {
		http.Error(w, fmt.Sprintf("Validation error: %s", message), http.StatusBadRequest)
		return
	}

	// Validation passed, we can now create the account here

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func removeFormat(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
}

func validateUser(u User) string {
	values := reflect.ValueOf(u)
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).String()
		//Check data field has been filled for all values
		if v == "" {
			return "data field is empty"
		}
	}
	//Validate email has an @ in middle
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return "email is invalid"
	}
	//Validate phone number is 10 digit int
	u.Phone = removeFormat(u.Phone)
	if _, err := strconv.Atoi(u.Phone); err != nil || len(u.Phone) != 10 {
		return "phone number is invalid"
	}
	//Validate skill level is an at least 0
	if u.SkillLevel < 0 {
		return "skill level is negative"
	}
	return ""
}
