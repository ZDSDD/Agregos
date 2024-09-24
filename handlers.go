package main

import (
	"blogAgg/internal/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login with no user specified")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("User with name %s doesn't exists. Check again if username is spelled correctly.", cmd.args[0])
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User successfully logged\n")
	return nil
}

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register with no user specified")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("user with name %s already exists", cmd.args[0])
	}
	log.Printf("Processing to create an user with name %s", cmd.args[0])
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}
	fmt.Printf("User %s successfully registered!\n", cmd.args[0])
	s.cfg.SetUser(user.Name)
	return nil
}
