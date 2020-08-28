package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/repo"
)

func main() {
	db := repo.Connect("development")
	defer db.Close()

	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Favorite{})
	db.DropTableIfExists(&models.Memo{})
}