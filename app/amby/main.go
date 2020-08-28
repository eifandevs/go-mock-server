package main

import (
    "github.com/eifandevs/amby/router"
)

func main() {
  e := router.Init()
  e.Logger.Fatal(e.Start(":8080"))
}