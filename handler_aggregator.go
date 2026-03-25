package main

import (
	"context"
	"fmt"
	"html"
	"time"

	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"
	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/gatorapi"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("time between requests is required\n")
	}
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error parsing time between requests: %v\n", err)
	}
	if time_between_reqs <= 0 {
		return fmt.Errorf("time_between_reqs must be greater than 0")
	}
	fmt.Printf("Collecting feeds every %.1f minutes...\n", time_between_reqs.Minutes())
	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return fmt.Errorf("Error scraping feeds: %v\n", err)
		}
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

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %v\n", err)
	}
	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v\n", err)
	}
	post, err := gatorapi.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v\n", err)
	}
	post = unescapeFeed(post)
	for _, item := range post.Channel.Item {
		post_params := createPostFromFeedItem(item, feed.ID)
		_, err = s.db.CreatePost(context.Background(), post_params)
		if pqErr := pq.As(err, pqerror.UniqueViolation); pqErr != nil {
			continue
		}
		if err != nil {
			return fmt.Errorf("error creating post: %v\n", err)
		}
	}
	return nil
}

func createPostFromFeedItem(item gatorapi.RSSItem, feedID uuid.UUID) database.CreatePostParams {
	publishedAt, _ := time.Parse(time.RFC1123Z, item.PubDate)
	return database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       item.Title,
		Description: item.Description,
		Url:         item.Link,
		PublishedAt: publishedAt,
		FeedID:      feedID,
	}
}
