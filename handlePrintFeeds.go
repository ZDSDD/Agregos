package main

import (
	"context"
	"fmt"
)

func handlePrintFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("\nPRINTING FEEDS\n")
	for _, feed := range feeds {
		fmt.Printf("feedName: %s\nuser: %s\nurl: %s\n\n", feed.FeedName, feed.Username, feed.Url)
	}
	return nil
}
