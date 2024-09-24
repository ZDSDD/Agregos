package main

import (
	"context"
	"fmt"
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
