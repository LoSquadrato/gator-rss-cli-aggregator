package main

import (
	"context"
	"fmt"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %v", err)
	}
	if len(follows) == 0 {
		fmt.Printf("User %s is not following any feeds\n", user.Name)
		return nil
	}
	for _, follow := range follows {
		fmt.Printf("Feed: %s\n", follow.FeedName)
	}
	return nil
}
