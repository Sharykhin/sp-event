package view

import "time"

type (
	Event struct {
		Title string
		Date  time.Time
	}

	Consumer struct {
		Email     string
		Events    []Event
		LastLogin time.Time
	}
)
