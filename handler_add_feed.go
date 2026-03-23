package main

import (
	"context"
	"fmt"
	"time"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("feed name and URL are required\n")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v\n", err)
	}
	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding feed: %v\n", err)
	}
	err = handlerFollow(s, command{
		name: "follow",
		arguments: []string{
			url,
		},
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v\n", err)
	}
	fmt.Printf("Successfully added feed: %s (%s)\n", feed.Name, feed.Url)
	return nil
}
