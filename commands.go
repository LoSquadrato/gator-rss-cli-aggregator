// commands.go defines the command struct and the commands struct which holds a map of command names to their handlers.
// It also defines the run and register methods for the commands struct.
package main

import (
	"errors"
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commands[cmd.name]
	if !ok {
		fmt.Println(cmd.name)
		return errors.New("unknown command\n")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.commands == nil {
		c.commands = make(map[string]func(*state, command) error)
	}
	c.commands[name] = f
}
