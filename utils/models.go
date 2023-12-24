package utils

import (
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ApiKey string `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt.String(),
		UpdatedAt: dbUser.UpdatedAt.String(),
		ApiKey: dbUser.ApiKey,
	}
}
type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt.String(),
		UpdatedAt: dbFeed.UpdatedAt.String(),
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}
func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}
	return feeds
}
type FeedFollow struct {
	ID uuid.UUID `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}
func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt.String(),
		UpdatedAt: dbFeedFollow.UpdatedAt.String(),
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}
func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}
type Post struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
	Description string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url string `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
}
func DatabasePosToPost(dbPost database.Post) Post {
	description := ""
	if dbPost.Description.Valid {
		description = dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}
func DatabasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, DatabasePosToPost(dbPost))
	}
	return posts
}