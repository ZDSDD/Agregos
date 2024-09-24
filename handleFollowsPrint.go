package main

import (
	"context"
	"fmt"
)

func handleFollowingPrint(s *state, _ command) error {
	feeds, err := s.db.GetFeedFollow(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting feed follow: %v\n", err)
	}
	fmt.Printf("%s is following feeds: \n", s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Println(feed.Name)
	}
	return nil
}
