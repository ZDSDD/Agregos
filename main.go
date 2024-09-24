package main

import (
	"blogAgg/internal/config"
	"blogAgg/internal/database"
	"context"
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
	c.register("reset", handleReset)
	c.register("users", middlewareLoggedIn(handlePrintUsers))
	c.register("agg", handleAgg)
	c.register("addfeed", middlewareLoggedIn(handleCreateFeed))
	c.register("feeds", handlePrintFeeds)
	c.register("follow", middlewareLoggedIn(handleFollowFeed))
	c.register("following", handleFollowingPrint)
	c.register("unfollow", middlewareLoggedIn(handleUnfollow))

	err = c.cmds[cmd.name](s, cmd)
	if err != nil {
		log.Fatalf("Error processing %s with error: %v\n", cmd.name, err)
		os.Exit(1)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
