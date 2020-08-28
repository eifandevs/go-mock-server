package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/repo"
)

func main() {
  db := repo.Connect("development")
  defer db.Close()

  // create user table
  db.AutoMigrate(&models.User{})

  // create favorite table
  db.AutoMigrate(&models.Favorite{})

  // create memo table
  db.AutoMigrate(&models.Memo{})
}