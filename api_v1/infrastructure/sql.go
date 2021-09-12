package sql

import (
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var onceDb sync.Once

var instance *gorm.DB

func CONN() *gorm.DB {
	onceDb.Do(func() {
		db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}
		instance = db
	})
	return instance
}
