package main

import (
	"blogAgg/internal/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func handleCreateFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Create feed hanlder need 2 arguments: name, url. Provided: %d args", len(cmd.args))
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}
	log.Printf("Successfuly added feed\n")
	_, err = followFeed(s, user.ID, feed.ID)

	if err != nil {
		return nil
	}
	log.Printf("%s successfully followed the feed %s\n", s.cfg.CurrentUserName, feed.Name)
	return nil
}
