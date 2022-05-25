package migration

import (
	"fiber-joglo-dev/database"
	"fiber-joglo-dev/models/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migration Succesfully")
}
