package main

import (
	"github.com/eifandevs/gms/models"
	"github.com/eifandevs/gms/repo"
)

func main() {
	db := repo.Connect("development")
	defer db.Close()

	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Favorite{})
	db.DropTableIfExists(&models.Memo{})
}