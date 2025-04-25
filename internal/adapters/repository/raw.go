package repository

import "time"

type RawPerson struct {
	ID          int
	Name        string
	Surname     string
	Patronymic  *string
	Age         *int
	Gender      *string
	Nationality *string
	CreatedAt   time.Time
}
