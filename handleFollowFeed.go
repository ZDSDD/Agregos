package main

import (
	"blogAgg/internal/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func handleFollowFeed(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Follow feed requires 1 argument. Provided: %d, %v", len(cmd.args), cmd.args)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error when getting an user: %v\n", err)
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error when getting a feed %s: %v\n", cmd.args[0], err)
	}

	_, err = followFeed(s, user.ID, feed.ID)

	if err != nil {
		return nil
	}
	log.Printf("%s successfully followed the feed %s\n", s.cfg.CurrentUserName, feed.Name)
	return nil
}

func followFeed(s *state, userId, feedId uuid.UUID) (database.CreateFeedFollowRow, error) {
	return s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    userId,
			FeedID:    feedId,
		})
}
