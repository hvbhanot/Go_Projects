package models

import (
	"RestAPI/db"
	"time"
)

// Event represents an event with its properties and methods.
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// events is a slice of Event structs, used to store a collection of events.
var events = []Event{}

// Save saves the event to the database. It inserts a new record into the "events" table,
// with the event's name, description, location, datetime, and user_id as values.
// It returns an error if there is an issue with the database query or execution.
// The last inserted ID is retrieved and assigned to the event's ID field.
func (event *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id
	return err
}

// GetAllEvents retrieves all events from the database and returns them as a slice of Event structs.
// It executes the SQL query "SELECT * FROM events" and scans the results into Event objects.
// If an error occurs during the database query or scanning process, it returns nil and the error.
// Otherwise, it returns the slice of events and nil error.
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// GetEventByID retrieves an event from the database based on the provided event ID.
// It executes a SQL SELECT query and scans the retrieved row into an Event struct.
// If the event is found, it is returned along with nil error. If no event is found,
// or an error occurs during the fetching process, nil event and the error are returned.
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates the event details in the events table.
// It updates the name, description, location, and dateTime fields for the event with the specified ID.
// It returns an error if the update operation fails.
// The function prepares an SQL UPDATE statement to update the event details.
// It uses the global variable db.DB of type *sql.DB to prepare the statement.
// Once prepared, the statement is closed at the end of the function.
// Then, it executes the prepared statement and passes the updated event details as arguments.
// It returns the result of the Exec method, which is the number of affected rows and an error if any.
// The error is returned and can be handled by the calling code accordingly.
func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

// Delete removes the event from the database using the event's ID.
// It executes the SQL query "DELETE FROM events WHERE id = ?" on the global DB connection.
// Returns an error if the deletion operation fails.
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

// Register inserts a new registration record in the database for the given event and user ID.
// It executes an SQL query to insert the event ID and user ID into the "registrations" table.
// Returns an error if there was an issue preparing the SQL statement or executing the query.
func (event Event) Register(userId int64) error {
	query := "INSERT INTO registrations(eventId, userId) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

// CancelRegistration deletes a registration record from the "registrations" table
// for a specific event and user. It takes a userId as a parameter and uses the
// eventId from the Event struct on which it is called.
// It prepares a SQL DELETE statement with placeholders for the eventId and userId values.
// It then executes the statement with the eventId and userId as arguments, deleting the
// registration record from the database.
// Returns an error if there was any issue with preparing the statement or executing the query.
func (event Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE eventId = ? AND userId = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)
	if err != nil {
		return err
	}
	return nil
}
