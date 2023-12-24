package main

import (
	"log"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting scraping with concurrency %d and time between requests %v", concurrency, timeBetweenRequest)
/* 	TODO: keep watching video https://youtu.be/un6ZyFkqFKo?t=31941
 */
}