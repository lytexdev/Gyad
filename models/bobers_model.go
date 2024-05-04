package models

import (
	"time"
)

type Bobers struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
