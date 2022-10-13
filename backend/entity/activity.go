package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Aname string
	// Email           string `gorm:"uniqueIndex"`
	Aemail    string
	Apassword string
	Activitys []Activity `grom:"foreignkey:AdminID"`
}

type Teacher struct {
	gorm.Model
	TfirstName string
	TlastName  string
	Temail     string `gorm:"uniqueIndex"`
	Tcontact   string
	Activitys  []Activity `grom:"foreignkey:TeacherID"`
}

type Location struct {
	gorm.Model
	Lname     string
	Activitys []Activity `grom:"foreignkey:LocationID"`
}

type Activity struct {
	gorm.Model
	Acname     string
	date_s     time.Time
	date_e     time.Time
	time_s     time.Time
	time_e     time.Time
	TeacherID  *uint
	Teacher    Teacher
	LocationID *uint
	Location   Location
	AdminID    *uint
	Admin      Admin
}
