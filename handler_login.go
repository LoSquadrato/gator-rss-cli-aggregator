// handler_login.go implements the login command for the Gator RSS CLI Aggregator.
package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("username is required\n")
	}
	name := cmd.arguments[0]
	// Check if user exists
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	switch {
	case err == sql.ErrNoRows:
		return fmt.Errorf("no user with id %s\n", name)
	case err != nil:
		return fmt.Errorf("query error: %v\n", err)
	default:
		// Set user in config
		if err := s.cfg.SetUser(name); err != nil {
			return fmt.Errorf("error setting user: %v\n", err)
		}
	}
	fmt.Printf("Successfully set user to %s\n", name)
	return nil
}
