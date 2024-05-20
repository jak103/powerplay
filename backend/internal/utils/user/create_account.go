package user

import (
	"encoding/json"
	"net/http"
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

	validateUser(user)

	// Validation passed, we can now create the account here

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func validateUser(u User) {
}
