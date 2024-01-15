package migrations

import (
	"fmt"
	"log"

	"github.com/ddfrmnsh-dev/mp-go-motivation/models"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils"
)

func Migration() {
	err := utils.DB.AutoMigrate(
		&models.Motivation{},
	)
	if err != nil {
		log.Fatal("Failed to migrate")
	}

	fmt.Println("Migrated successfully")
}
