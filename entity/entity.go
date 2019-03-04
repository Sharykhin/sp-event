package entity

import "time"

type (
	User struct {
		Email string
	}

	Event struct {
		Title string
		Date time.Time
	}

	Producer struct {
		Email string
	}

	Consumer struct {
		Email string
	}
)