package main

import (
	"context"
	"fmt"
	"time"

	"github.com/K-Road/rss_feed_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	url := cmd.Arguments[0]
	now := time.Now().UTC()
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error: unable to add feedfollow: %w", err)
	}

	fmt.Printf("Feed follow created successfully:")
	printFeedFollow(feedfollow.UserName, feedfollow.FeedName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(follows) == 0 {
		fmt.Println("No follows")
		return nil
	}

	for _, follow := range follows {
		fmt.Printf("* Name:	%s\n", follow.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
