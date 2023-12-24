package utils

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/google/uuid"
)

func StartScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting on %v gorotuines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println(err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(errors.New("Error marking fetch as fetched: " + err.Error()))
	}

	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Println(errors.New("Error fetching feed: " + err.Error()))
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		published_date, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println(errors.New("Error parsing published date: " + err.Error()))
			continue
		}

		_, err = db.CreatePost(context.Background(),
		database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Description: description,
			PublishedAt: published_date,
			Url: item.Link,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println(errors.New("Error creating post: " + err.Error()))
			return
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Url, len(rssFeed.Channel.Item))
}