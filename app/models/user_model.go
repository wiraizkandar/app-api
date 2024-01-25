package models

import (
	"time"
)

type User struct {
	Id              string    `json:"-"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"-"`
	Password        string    `json:"-"`
	RememberToken   bool      `json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
