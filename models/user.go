package models

import (
	"github.com/GymWorkoutApp/gwap-auth/utils/uuid"
	"github.com/jinzhu/gorm"
)

// Client client model
type User struct {
	Base
	ID     		string `json:"id" gorm:"not null;primary_key;"`
	Username 	string `json:"username" validate:"required" gorm:"not null;"`
	Password 	string `json:"password" validate:"required" gorm:"not null;"`
	Name 		string `json:"name" validate:"required" gorm:"not null;"`
}

func (u User) GetID() string {
	return u.ID
}

func (u User) SetID(id string) {
	u.ID = id
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) SetUsername(username string) {
	u.Username = username
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) SetPassword(password string) {
	u.Password = password
}

func (u User) GetName() string {
	return u.Name
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.Must(uuid.NewRandom()).String())
	return nil
}