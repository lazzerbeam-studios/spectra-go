package models

import "time"

type Profile struct {
	ID int `json:"id"`

	Email string `json:"email"`

	Name  string `json:"name"`
	Image string `json:"image"`

	Created time.Time `json:"created"`
}

type ProfileUpdate struct {
	_ struct{} `json:"-" additionalProperties:"true"`

	ID int `json:"id"`

	Name string `json:"name"`
}

type ProfileUpdateEmail struct {
	_ struct{} `json:"-" additionalProperties:"true"`

	ID int `json:"id"`

	Email string `json:"email"`
}

type User struct {
	ID int `json:"id"`

	Email string `json:"email"`

	Name  string `json:"name"`
	Image string `json:"image"`
}
