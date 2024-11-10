package domain

import "time"

type Author struct {
	ID        int       `db:"id"`
	Fullname  string    `db:"fullname"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
