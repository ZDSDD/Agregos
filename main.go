package main

import (
	"blogAgg/internal/config"
	"blogAgg/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	var cfg, err = config.Read()
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Succesffuly connected to the database!")
	var s = &state{
		cfg: cfg,
	}

	s.db = dbQueries
	var c commands = commands{
		cmds: make(map[string]func(*state, command) error),
	}

	if len(os.Args) < 2 {
		log.Fatal("You need to specify command")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	c.register("login", handleLogin)
	c.register("register", registerHandler)
	c.register("reset", handleReset)
	c.register("users", middlewareLoggedIn(handlePrintUsers))
	c.register("agg", handleAgg)
	c.register("addfeed", middlewareLoggedIn(handleCreateFeed))
	c.register("feeds", handlePrintFeeds)
	c.register("follow", middlewareLoggedIn(handleFollowFeed))
	c.register("following", handleFollowingPrint)
	c.register("unfollow", middlewareLoggedIn(handleUnfollow))

	_, ok := c.cmds[cmd.name]
	if !ok {
		log.Fatalf("%s command not known.\n", cmd.name)
	}
	err = c.cmds[cmd.name](s, cmd)
	if err != nil {
		log.Fatalf("Error processing %s with error: %v\n", cmd.name, err)
		os.Exit(1)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}

func scrapFeeds(s *state) error {
	feed_to_fetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), feed_to_fetch.ID)
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), feed_to_fetch.Url)
	if err != nil {
		return err
	}

	fmt.Printf("\n\nPrinting items for %s\n\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		time.Sleep(time.Millisecond * 33)
		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: StringToNullString(item.Description),
			PublishedAt: DateToNullDate(item.PubDate),
			FeedID:      feed_to_fetch.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	return nil
}

// Helper function to convert a string to sql.NullString
func StringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func DateToNullDate(d string) sql.NullTime {
	if d == "" {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}
	layout := "2006-01-02" // Adjust this to match your date format
	parsedDate, err := time.Parse(layout, d)
	if err != nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: parsedDate, Valid: true}
}
