// task-profile-service/database/connection.go

package database

import (
	"fmt"
	"task-profile-service/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	fmt.Println(config.Cnfg.Database)
	db, err := gorm.Open("mysql", config.Cnfg.Database.Username+":"+config.Cnfg.Database.Password+"@tcp("+config.Cnfg.Database.Host+":"+config.Cnfg.Database.Port+")/"+config.Cnfg.Database.Name+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("connected to db")
	DB = db
}
