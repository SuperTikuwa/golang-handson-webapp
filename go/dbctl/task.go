package dbctl

import (
	"log"

	"github.com/SuperTikuwa/golang-handson-todo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	// user:pass@tcp(127.0.0.1:3306)/dbname
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:33063)/todo_db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertNewTask(task model.Task) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}
