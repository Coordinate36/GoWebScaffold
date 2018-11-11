package models

import (
	. "web/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Model struct {
}

var Db *gorm.DB

func (m *Model) Setup() {
	var err error
	Db, err = gorm.Open(Config.DbDriver, Config.DbUri)
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Person{})
}

func (m *Model) Close() {
	Db.Close()
}
