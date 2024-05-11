package schedule

import (
	"database/sql"
)

type IceTime struct {
	ID        int    `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TeamSchedule struct {
	ID       int       `json:"id"`
	TeamID   int       `json:"team_id"`
	TeamName string    `json:"team_name"`
	IceTimes []IceTime `json:"ice_times"`
}

type LeagueSchedule struct {
	ID            int            `json:"id"`
	LeagueID      int            `json:"league_id"`
	LeagueName    string         `json:"league_name"`
	TeamSchedules []TeamSchedule `json:"team_schedules"`
}

type SeasonSchedule struct {
	ID              int              `json:"id"`
	SeasonID        int              `json:"season_id"`
	SeasonName      string           `json:"season_name"`
	LeagueSchedules []LeagueSchedule `json:"league_schedules"`
}

func getIceTimes(db *sql.DB) ([]IceTime, error) {
	query := `SELECT * FROM ice_time`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
			return
		}
	}(rows)

	var iceTimes []IceTime
	for rows.Next() {
		var iceTime IceTime
		err := rows.Scan(&iceTime.ID, &iceTime.StartTime, &iceTime.EndTime)
		if err != nil {
			return nil, err
		}
		iceTimes = append(iceTimes, iceTime)
	}
	return iceTimes, nil
}

func createSeasonScheduleTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS season_schedule (
    		id SERIAL PRIMARY KEY,
    		season_id INT NOT NULL,
    		season_name TEXT NOT NULL,
    		league_schedules JSONB NOT NULL
    )`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func insertSeasonSchedule(db *sql.DB, seasonSchedule SeasonSchedule) (int, error) {
	query := `INSERT INTO season_schedule (season_id, season_name, league_schedules) VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(query, seasonSchedule.SeasonID, seasonSchedule.SeasonName, seasonSchedule.LeagueSchedules).Scan(&pk)
	if err != nil {
		return 0, err
	}
	return pk, nil
}
