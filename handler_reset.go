package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: unable to delete: %w", err)
	}
	fmt.Println("Database reset successfully")
	return nil
}
