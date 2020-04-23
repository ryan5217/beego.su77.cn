package models

import (
	"time"
)

type UserProfile struct {
	Id int `json:"id"`
	Sex bool `json:"sex"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
	Users *Users `json:"users" orm:"rel(fk)"`
}

