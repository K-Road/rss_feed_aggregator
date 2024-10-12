package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("error: login command requires 1 argument, the username")
	}
	err := s.cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("%v has been set.\n", cmd.Arguments[0])
	return nil
}
