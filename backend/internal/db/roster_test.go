package db

import (
	"time"

	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) TestCreateRoster() {
	layout := "2006-01-02 15:04:05"
	value := "2023-12-14 17:09:47"

	date, _ := time.Parse(layout, value)

	// Create the Captain of the Roster
	captain := models.User{
		FirstName:    "John",
		LastName:     "Smith",
		Email:        "test@email.com",
		Password:     "password",
		Phone:        "7",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	dbCapt, err := s.session.CreateUser(&captain)
	s.Nil(err)
	s.NotNil(dbCapt)

	capt, err := s.session.GetUserByID(dbCapt.ID)
	s.Nil(err)
	s.NotNil(capt)

	// Create the list of Players
	user1 := models.User{
		FirstName:    "Jack",
		LastName:     "Smith",
		Email:        "test1@email.com",
		Password:     "password",
		Phone:        "8",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	user2 := models.User{
		FirstName:    "Jill",
		LastName:     "Smith",
		Email:        "test2@email.com",
		Password:     "password",
		Phone:        "9",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	dbUser1, err := s.session.CreateUser(&user1)
	s.Nil(err)
	s.NotNil(dbUser1)

	dbUser2, err := s.session.CreateUser(&user2)
	s.Nil(err)
	s.NotNil(dbUser2)

	playerIDs := []uint{dbUser1.ID, dbUser2.ID}

	players, err := s.session.GetUsersByIDs(playerIDs)
	s.Nil(err)
	s.NotNil(players)

	// Create the Roster
	roster := models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		Players:   players,
	}

	myRoster, err := s.session.CreateRoster(&roster)
	s.Nil(err)
	s.NotNil(myRoster)

	// Verify the number of players on the team besides the Captain
	// equals 2
	s.Len(myRoster.Players, 2)

	s.Equal(uint(1), myRoster.ID)

	// Test Getting Rosters
	rosters, err := s.session.GetRosters()
	s.Nil(err)
	s.Len(rosters, 1)
}
