package main

import (
	"context"
	"fmt"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("feed url is required\n")
	}
	url := cmd.arguments[0]
	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v\n", err)
	}
	_, err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v\n", err)
	}
	fmt.Printf("Successfully unfollowed feed: %s (%s)\n", feed.Name, feed.Url)
	return nil
}
