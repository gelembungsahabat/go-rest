package models

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"type:uint;primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Nickname  string    `json:"nickname" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "testdb"
)

var DB *gorm.DB

func ConnectDatabase() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic("failed to connect db!")
	}
	db.AutoMigrate(&User{})
	DB = db
}
