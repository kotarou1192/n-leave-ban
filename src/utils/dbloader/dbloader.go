package dbloader

import (
	"../envloader"

	"github.com/jinzhu/gorm"
)

// InitDB load structure from db
func InitDB() *gorm.DB {
	conf := envloader.LoadConfig()
	db, err := gorm.Open("sqlite3", "db/"+conf.ENV+".db")
	if err != nil {
		panic(err)
	}

	return db
}
