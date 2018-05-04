package utils

import (
	"github.com/jinzhu/gorm"
)

const (
	DB_HOST = "106.14.2.153"
	DB_PORT = "3306"
	DB_USERNAME = "root"
	DB_PASSWORD = "root811123"
	DB_NAME = "inj_zone"
)

func OpenConnection() (db *gorm.DB, err error) {
	conn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", conn)
	return db, err
}
