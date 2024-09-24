package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	err := s.db.ClearTable(context.Background())
	if err != nil {
		return fmt.Errorf("Error on table clear: %v\n", err)
	}
	fmt.Println("Table cleared successfully")
	return nil
}
