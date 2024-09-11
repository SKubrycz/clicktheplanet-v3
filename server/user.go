package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint64    `json:"id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(login, email, password string) (*User, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Login:     login,
		Email:     email,
		Password:  string(encrypted),
		CreatedAt: time.Now().UTC(),
	}, nil
}
