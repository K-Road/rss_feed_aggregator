package main

import (
	"context"
	"fmt"
	"time"

	"github.com/K-Road/rss_feed_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("error: Couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error: Couldn't set current users: %w", err)
	}

	fmt.Printf("%v has been set.\n", cmd.Arguments[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Arguments[0]
	now := time.Now().UTC()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error: Can't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error: Couldn't set current users: %w", err)
	}

	fmt.Printf("%s created.", name)
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)

}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: unable to delete: %w", err)
	}
	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user)
		} else {
			fmt.Printf("* %v\n", user)
		}
	}
	return nil
}
