package main

import (
	"blogAgg/internal/config"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello world")
	var cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("username: %s\nconnection: %s\n", cfg.CurrentUserName, cfg.DbURL)
	cfg.SetUser("Sebasadowy")
}
