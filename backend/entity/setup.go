package entity

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Admin{},
		&Teacher{},
		&Location{},
		&Activity{},
	)

	db = database

	// -------SETUP DATA-------
	// Admin
	thanadet := Admin{
		Aname:     "Thanadet Choedaimueang",
		Aemail:    "tun@gmail.com",
		Apassword: "asdfg123456",
	}
	theerawat := Admin{
		Aname:     "Theerawat Goodking",
		Aemail:    "theerawatG@gmail.com",
		Apassword: "123456",
	}
	db.Model(&Admin{}).Create(&thanadet)
	db.Model(&Admin{}).Create(&theerawat)

	Pornsak := Teacher{
		TfirstName: "Pornsak",
		TlastName:  "Songsang",
		Temail:     "pornsakS@gmail.com",
		Tcontact:   "111-222-3333",
	}
	db.Model(&Teacher{}).Create(&Pornsak)

	// Location
	L1 := Location{
		Lname: "ลานกิจกรรม",
	}
	L2 := Location{
		Lname: "หอประชุม",
	}
	db.Model(&Location{}).Create(&L1)
	db.Model(&Location{}).Create(&L2)
	// Activity
	Ac1 := Activity{
		Acname:   "Activity festival",
		date_s:   time.Now(),
		date_e:   time.Now(),
		time_s:   time.Now(),
		time_e:   time.Now(),
		Location: L1,
		Teacher:  Pornsak,
		Admin:    theerawat,
	}
	db.Model(&Activity{}).Create(&Ac1)

	Ac2 := Activity{
		Acname:   "Freshman 2022",
		date_s:   time.Now(),
		date_e:   time.Now(),
		time_s:   time.Now(),
		time_e:   time.Now(),
		Location: L1,
		Teacher:  Pornsak,
		Admin:    theerawat,
	}
	db.Model(&Activity{}).Create(&Ac2)
	fmt.Printf("\nEnd Querry\n")
}
