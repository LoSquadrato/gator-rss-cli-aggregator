package main

import (
	"context"
	"fmt"
	"html"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/gatorapi"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	cmd.arguments = append(cmd.arguments, feedURL)
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("feed URL is required\n")
	}
	feedURL := cmd.arguments[0]
	feed, err := gatorapi.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v\n", err)
	}
	// unescape HTML entities in title and description
	feed = unescapeFeed(feed)
	fmt.Printf("Title: %s\nDescription: %s\nLink: %s\n", feed.Channel.Title, feed.Channel.Description, feed.Channel.Link)
	for _, item := range feed.Channel.Item {
		fmt.Printf("Title: %s\nDescription: %s\nLink: %s\nPubDate: %s\n", item.Title, item.Description, item.Link, item.PubDate)
	}
	return nil
}

func unescapeFeed(feed *gatorapi.RSSFeed) *gatorapi.RSSFeed {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
	return feed
}
