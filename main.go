package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

type ShowID int

const (
	OnePieceShowID ShowID = 347

	LatestDownloadedEpisodeFilename = ".latest_downloaded_episode"
)

func LatestDownloadedEpisode(dataDir string) (int, error) {
	file := path.Join(dataDir, LatestDownloadedEpisodeFilename)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, fmt.Errorf("file could not be read %v: %w", file, err)
	}
	episodeNumber, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("content of %v does not look like an number: %w", file, err)
	}
	return episodeNumber, nil
}

func UpdateLatestDownloadedEpisode(dataDir string, episodeNumber int) error {
	file := path.Join(dataDir, LatestDownloadedEpisodeFilename)
	err := ioutil.WriteFile(file, []byte(strconv.Itoa(episodeNumber)), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Fatalf("[ERROR] Environ variable SENDGRID_API_KEY must be set")
	}

	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		log.Fatalf("[ERROR] Environ variable DATA_DIR must be set")
	}

	scrapeSchedule := os.Getenv("SCRAPE_SCHEDULE")
	if scrapeSchedule == "" {
		log.Fatalf("[ERROR] Environment variable SCRAPE_SCHEDULE must be set")
	}

	notificationEmail := os.Getenv("NOTIFICATION_EMAIL")

	schedule, err := cron.ParseStandard(scrapeSchedule)
	if err != nil {
		log.Fatalf("[ERROR] Failed to parse SCRAPE_SCHEDULE as a standard cron specification: %v", err)
	}

	for {
		nextRun := schedule.Next(time.Now())
		log.Printf("[INFO] Next try at %v", nextRun)
		<-time.After(nextRun.Sub(time.Now()))

		latestDownloadedEpisode, err := LatestDownloadedEpisode(dataDir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				latestDownloadedEpisode = 890
			} else {
				log.Printf("[ERROR] Could not get latest downloaded episode number: %v", err)
				continue
			}
		}

		// try to download next episode
		nextEpisode := latestDownloadedEpisode + 1
		log.Printf("[INFO] Download episode %v", nextEpisode)

		if err := DownloadEpisode(dataDir, OnePieceShowID, nextEpisode); err != nil {
			if errors.Is(err, ErrEpisodeNotFound) {
				log.Printf("[INFO] Episode %v not found", nextEpisode)
			} else {
				log.Printf("[ERROR] Failed to download episode %v: %v", nextEpisode, err)
			}
			continue
		}
		if err := UpdateLatestDownloadedEpisode(dataDir, nextEpisode); err != nil {
			log.Printf("[ERROR] Failed to update latest episode %v: %v", nextEpisode, err)
		}

		subject := fmt.Sprintf("[Horrible Crawler] New episode of One Piece %v available for watching", nextEpisode)
		content := fmt.Sprintf("One Piece episode %v successfully downloaded", nextEpisode)

		log.Printf("[INFO] %v", subject)

		if notificationEmail != "" {
			if err := SendNotification(apiKey, notificationEmail, subject, content); err != nil {
				fmt.Printf("[ERROR] failed to send notification: %v\n", err)
				continue
			}
		}
	}
}
