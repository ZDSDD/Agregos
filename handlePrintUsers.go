package main

import (
	"blogAgg/internal/database"
	"context"
	"fmt"
)

func handlePrintUsers(s *state, _ command, user database.User) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error on db.GetUsers: %v\n", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}
