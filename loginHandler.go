package main

import "fmt"

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login with no user specified")
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User successfully logged\n")
	return nil
}
