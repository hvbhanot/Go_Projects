package Modals

import (
	"RestAPI/db"
	"database/sql"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e *Event) SaveNewEvent() error {
	query := `INSERT INTO events(
                   name,
                   description,
                   location,
                   dateTime,	
                   userID )
                 
                 VALUES (?,?,?,?,?)   
                   
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := stmt.Close()
		if closeErr != nil {
			log.Printf("Error closing statement: %v", closeErr)
		}
	}()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
