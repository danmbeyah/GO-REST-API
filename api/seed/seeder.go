
package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"go-rest/api/models"
)

var users = []models.User{
	models.User{
		Username: "user1",
		Email:    "user1@example.com",
		Password: "password1",
	},
	models.User{
		Username: "user2",
		Email:    "user2@example.com",
		Password: "password2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}