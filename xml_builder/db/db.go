package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("E:/1/Govno/utm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Table{},
	)

	DB = database
}

type Table struct {
	Amc     string `gorm:"primarykey"`
	F2RegId string
	Volume  string
}
