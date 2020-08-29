package main

import (
	"github.com/eifandevs/main/models"
	"github.com/eifandevs/main/repo"
)

func main() {
	db := repo.Connect("development")
	defer db.Close()

	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Favorite{})
	db.DropTableIfExists(&models.Memo{})
}