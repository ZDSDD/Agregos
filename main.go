package main

import (
	"blogAgg/internal/config"
	"blogAgg/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	var cfg, err = config.Read()
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Succesffuly connected to the database!")
	var s = &state{
		cfg: cfg,
	}

	s.db = dbQueries
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
	c.register("register", registerHandler)

	err = c.cmds[cmd.name](s, cmd)
	if err != nil {
		log.Fatalf("Error processing command with error: %v\n", err)
		os.Exit(1)
	}

}
