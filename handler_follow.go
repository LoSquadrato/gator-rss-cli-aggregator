package main

import (
	"context"
	"fmt"
	"time"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	url := cmd.arguments[0]
	if url == "" {
		return fmt.Errorf("error missing feed url")
	}
	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed from url: %v", err)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Printf("Successfully followed feed: %s (added by %s)\n", feedFollow.FeedName, feedFollow.UserName)
	return nil

}
