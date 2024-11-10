package domain

import "time"

type Blog struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Slug      string    `db:"slug"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
