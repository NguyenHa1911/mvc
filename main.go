package main

import (
	"fmt"

	"HA_MVC.Form/Database"
	"HA_MVC.Form/models"
	"HA_MVC.Form/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Database.DB, err = gorm.Open("mysql", Database.DbURL(Database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Database.DB.Close()
	Database.DB.AutoMigrate(&models.Form{})
	r := routes.SetupRouter()

	//running
	r.Run(":3000")
}
