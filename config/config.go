// task-profile-service/config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-profile-service/models"
)

var Cnfg models.Config
func LoadConfig(filePath string) {
	var config models.Config

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("json opened")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	Cnfg = config
}
