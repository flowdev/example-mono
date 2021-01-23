package entities

import "time"

type Money int64

type Product struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Price       Money
}
