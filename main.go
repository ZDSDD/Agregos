package main

import (
	"blogAgg/internal/config"
	"fmt"
	"log"
	"os"
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
	c.register("login", func(s *state, c command) error {
		if len(c.args) != 1 {
			return fmt.Errorf("login requires an username.")
		}
		s.cfg.SetUser(c.args[0])
		return nil
	})
	if err = c.cmds["login"](s, cmd); err != nil {
		os.Exit(1)
	}
}
