package main

import (
	"blogAgg/internal/config"
	"blogAgg/internal/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello world")
	var cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	var s = &state{
		cfg: cfg,
	}

	var c commands = commands{
		cmds: make(map[string]func(*state, command) error),
	}

	if len(os.Args) < 2 {
		log.Fatal("You need to specify command")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	c.register("login", handleLogin)
	if err = c.cmds["login"](s, cmd); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	s.db = dbQueries

}
