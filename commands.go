package main

import "errors"

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.registeredCommands == nil {
		c.registeredCommands = make(map[string]func(*state, command) error)
	}
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, exists := c.registeredCommands[cmd.Name]; exists {
		return f(s, cmd)
	}
	return errors.New("command not found: ")
}
