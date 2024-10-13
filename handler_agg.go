package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Arguments) != 1 {
	// 	return fmt.Errorf("usage: %v <name>", cmd.Name)
	// }

	//url := cmd.Arguments[0]
	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error: unable to fetch: %w", err)
	}

	fmt.Printf("%+v", rssFeed)

	return nil
}
