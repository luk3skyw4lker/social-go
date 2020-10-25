package models_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
	"github.com/luk3skyw4lker/social-go/api/controllers"
	"github.com/luk3skyw4lker/social-go/api/models"
)

var server controllers.Server
var userInstance models.User
var postInstance models.Post

func TestMain(m *testing.M) {
	var err error

	err = godotenv.Load(os.ExpandEnv("../../.env"))

	if err != nil {
		log.Fatalf("Error getting .env file %v", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("DB_DRIVER")

	switch TestDbDriver {
	case "mysql":
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			log.Fatalf("Error on opening databse %v", err)
		} else {
			fmt.Println("We are connected to the MySQL database")
		}

	case "postgres":
		DBURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

		fmt.Println(DBURL)
		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			log.Fatalf("Error on opening database: %v", err)
		} else {
			fmt.Println("We are connected to the Postgres database")
		}

	default:
		fmt.Printf("Unrecognized DB Driver: %s", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.Post{}, &models.User{}).Error

	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error

	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed users table")

	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()

	user := models.User{
		Nickname: "Fake",
		Email:    "fake@faker.com",
		Password: "fake@password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func seedUsers() error {
	users := []models.User{
		models.User{
			Nickname: "Fake 1",
			Email:    "fake1@faker.com",
			Password: "fake1@password",
		},
		models.User{
			Nickname: "Fake 2",
			Email:    "fake2@faker.com",
			Password: "fake2@password",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error

	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error

	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed users and posts tables")

	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {
	err := refreshUserAndPostTable()

	if err != nil {
		return models.Post{}, err
	}

	user := models.User{
		Nickname: "Fake 1",
		Email:    "fake1@faker.com",
		Password: "fake1@password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error

	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Title:    "Title",
		Content:  "This is the content",
		AuthorID: user.ID,
	}

	err = server.DB.Model(&models.Post{}).Create(&post).Error

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {
	var err error
	var users = []models.User{
		models.User{
			Nickname: "Fake 1",
			Email:    "fake1@faker.com",
			Password: "fake1@password",
		},
		models.User{
			Nickname: "Fake 2",
			Email:    "fake2@faker.com",
			Password: "fake2@password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "This is the content",
		},
		models.Post{
			Title:   "Title 2",
			Content: "This is the content",
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			return []models.User{}, []models.Post{}, err
		}

		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error

		if err != nil {
			return []models.User{}, []models.Post{}, err
		}
	}

	return users, posts, nil
}
