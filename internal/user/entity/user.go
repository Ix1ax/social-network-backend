package entity

import (
	"time"

	"github.com/google/uuid"
)

/*
 * Помимо типа данных написано как в базе данных и как в json оно отображается
 */
type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Surname      string    `db:"surname" json:"surname"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
