package seed

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/luk3skyw4lker/social-go/api/models"
)

var users = []models.User{
	{
		Nickname: "Lucas Henrique",
		Email:    "lucashenriqueblemos@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Another user",
		Email:    "another@email.com",
		Password: "another",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load is...
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error

	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error

	if err != nil {
		log.Fatalf("Cannot migrate: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error

	if err != nil {
		log.Fatalf("Cannot add FK: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		fmt.Printf("%v", posts[i].AuthorID)

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error

		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
