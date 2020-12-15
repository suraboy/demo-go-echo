package models

import (
	"time"
)

type Users struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	Name        string    `json:"name" xml:"name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" xml:"email"`
	Verify      string    `json:"verify" gorm:"type:enum('waiting','yes','no');default:'waiting'"`
	Mobile      string    `json:"mobile"`
	Type        string    `json:"type" gorm:"type:enum('owner','staff','other','admin','customer','brand-owner');default:'other'"`
	Pin         string    `json:"pin"`
	Status      string    `json:"status" gorm:"type:enum('active', 'inactive', 'ban');default:'inactive'"`
	UserGroupId int       `json:"user_group_id"`
	Gender      string    `json:"gender" gorm:"type:enum('male', 'female');default:'male'"`
	Birthday    time.Time `json:"birthday"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}