package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		cmd.arguments = append(cmd.arguments, "2")
	}
	limit, err := strconv.Atoi(cmd.arguments[0])
	if err == nil {
		fmt.Printf("limit=%d, type: %T\n", limit, limit)
	}
	posts, err := s.db.GetPostForUser(context.Background(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts for user: %v\n", err)
	}
	for _, post := range posts {
		println("title: " + post.Title)
		println("url: " + post.Url)
		println("description: " + post.Description)
		println("published at: " + post.PublishedAt.String())
		println()
		println()
	}
	return nil
}
