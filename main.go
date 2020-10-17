package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luk3skyw4lker/social-go/api"
)

func main() {
	api.Run()
}
