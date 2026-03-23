package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeed(context.Background())
	if err != nil {
		return fmt.Errorf("error listing feeds: %v\n", err)
	}
	for i, feed := range feeds {
		name, err := s.db.GetUserFeed(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("can't get user name: %v\n", err)
		}
		println("number: " + fmt.Sprint(i+1))
		println("feed name: " + feed.Name)
		println("feed url: " + feed.Url)
		println("added by: " + name)
	}
	return nil
}
