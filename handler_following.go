package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user_name := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), user_name)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %v", err)
	}
	if len(follows) == 0 {
		fmt.Printf("User %s is not following any feeds\n", user)
		return nil
	}
	for _, follow := range follows {
		fmt.Printf("Feed: %s\n", follow.FeedName)
	}
	return nil
}
