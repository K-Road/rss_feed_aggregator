package main

import (
	"context"
	"fmt"
	"time"

	"github.com/K-Road/rss_feed_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.Arguments[0]
	url := cmd.Arguments[1]
	now := time.Now().UTC()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error: unable to add feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("============================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:		%s\n", feed.ID)
	fmt.Printf("* Created:	%s\n", feed.CreatedAt)
	fmt.Printf("* Updated:	%s\n", feed.UpdatedAt)
	fmt.Printf("* Name:		%s\n", feed.Name)
	fmt.Printf("* URL:		%s\n", feed.Url)
	fmt.Printf("* UserID:	%s\n", feed.UserID)
}
