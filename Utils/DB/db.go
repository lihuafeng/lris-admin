package DB

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var maxIdle, maxOpen, maxLifetime = 10, 100, 30

var db *gorm.DB

func Connect(driveName string, dns string) (err error) {
	drive, err := gorm.Open(driveName, dns)
	if err != nil {
		err = errors.New("databases connect failed.")
	}
	drive.DB().SetMaxIdleConns(maxIdle)
	drive.DB().SetMaxOpenConns(maxOpen)
	drive.DB().SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
	db = drive
	return
}

func New() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
