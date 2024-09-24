package main

import (
	"blogAgg/internal/database"
	"context"
	"fmt"
)

func handleUnfollow(s *state, cmd command, user database.User) error {

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error when getting a feed %s: %v\n", cmd.args[0], err)
	}

	return s.db.RemoveFeedFollow(context.Background(), database.RemoveFeedFollowParams{
		Name: user.Name,
		Url:  feed.Url,
	})

}
