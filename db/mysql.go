package db

import (
	"fmt"
	m "ginlearn/models"
	"github.com/jinzhu/gorm"

)

var db *gorm.DB
var user *m.User
var job *m.Job

func InitDB() {
	fmt.Println("InitDB...")
	newConn()
	//createTables(newConn())
}

func newConn() *gorm.DB {
	db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("new conn...")

	return db
}

func createTables(*gorm.DB) {
	fmt.Println(!db.HasTable(user))
	if !db.HasTable(user) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(user).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(job) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(job).Error; err != nil {
			panic(err)
		}
	}
}

func InsertUser(u *m.User) {
	fmt.Println("Insert User...")
	db.Create(u)
}

func InsertJob(job *m.Job) {
	fmt.Println("Insert Job...")
	db.Create(job)
}
