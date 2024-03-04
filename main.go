package main

import (
	"location_program/config"
	"log"
)

func main() {
	config := config.Init()
	log.Println(config)
}
