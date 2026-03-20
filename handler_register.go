// handler_register.go contains the handler for the "register" command, which allows users to create a new account.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("user name is required\n")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err == nil {
		return fmt.Errorf("user with name %s already exists\n", cmd.arguments[0])
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v\n", err)
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("Successfully created user %s\n", user.Name)
	log.Printf("id: %s, created_at: %s, updated_at: %s, name: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
