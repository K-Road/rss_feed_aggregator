package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/K-Road/rss_feed_aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <time_between_requests>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("error fetching next feed: %v", err)
		return
	}

	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("error fetching feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("* Title:	%s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
