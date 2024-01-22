package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	FullName string    `db:"full_name"`
	Email    string    `db:"email"`
}
