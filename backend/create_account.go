package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName  string `json:"firstName"  validate:"nonzero"`
	LastName   string `json:"lastName"   validate:"nonzero"`
	Email      string `json:"email"      validate:"nonzero"`
	Password   string `json:"password"   validate:"nonzero"`
	Phone      string `json:"phone"      validate:"nonzero"`
	Birthday   string `json:"birthday"   validate:"nonzero"`
	Experience int    `json:"experience" validate:"min=0"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Parse the JSON request and populate the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate the User struct

	// Validation passed, proceed to process the user data
	// This is where we actually create the account

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully!")
}
