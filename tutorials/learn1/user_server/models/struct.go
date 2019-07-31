package models

import "time"

type User struct {
	Id           int        `json:"id"`
	UserId       int        `json:"UserId"`
	UserName     string     `json:"name"`
	Pwd          string     `json:"pwd"`
	UpdatedTime  time.Time  `json:"updated_at" xorm:"updated"`
	CreatedTime  time.Time  `json:"created_at" xorm:"created"`
}
