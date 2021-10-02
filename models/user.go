package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	Email     string     `json:"email" form:"email"`
	Address   string     `json:"address" form:"address"`
	Age       int        `json:"age" form:"age"`
	Height    float32    `json:"height" form:"height"`
	Public    bool       `json:"public" form:"public"`
	Birthday  *time.Time `json:"birthday"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
}
