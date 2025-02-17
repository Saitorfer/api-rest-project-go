package models

import (
	"time"

	"practise.com/rest-api-go/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	//use ? as a secure sintax (to avoid sql Injections)
	query := `INSERT INTO events
	(name, description, location,dateTime,user_id) 
	VALUES
	(?,?,?,?,?)
	`
	//first prepare
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	//close the execution with defer, so it always close
	defer stmt.Close()
	//execute
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	//for autoincrement of id
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"

	//just get 1 row
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

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
	//close the execution with defer, so it always close
	defer stmt.Close()

	//execute
	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.ID)

	return err

}

func (event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	//close the execution with defer, so it always close
	defer stmt.Close()

	//execute
	_, err = stmt.Exec(event.ID)

	return err
}
