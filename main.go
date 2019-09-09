package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/antchfx/htmlquery"
)

type ShowID int

var onePieceShowID ShowID = 347

func ShowURL(showID ShowID) string {
	return fmt.Sprintf(`https://horriblesubs.info/api.php?method=getshows&type=show&showid=%v`, showID)
}

var ErrEpisodeNotFound = fmt.Errorf("episode not found")

func getMagnet(episodeNumber int) (string, error) {
	showURL := ShowURL(onePieceShowID)
	node, err := htmlquery.LoadURL(showURL)
	if err != nil {
		return "", fmt.Errorf("failed to load url %q: %w", showURL, err)
	}

	episode := htmlquery.FindOne(node, fmt.Sprintf(`//*[@id="%v-1080p"]`, episodeNumber))
	if episode == nil {
		return "", fmt.Errorf("could not find magnet link for episode %v (1080p): %w", episodeNumber, ErrEpisodeNotFound)
	}

	link := htmlquery.FindOne(episode, `//*[@class="dl-type hs-magnet-link"]`)
	if link == nil {
		return "", fmt.Errorf("magnet link for episode %v (1080p) not found", episodeNumber)
	}

	href := htmlquery.SelectAttr(link.FirstChild, "href")
	if href == "" {
		return "", fmt.Errorf("href attribute for episode %v (1080p) is empty", episodeNumber)
	}

	return href, nil
}

const sendgridMailSendAPI = "https://api.sendgrid.com/v3/mail/send"

type MailContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MailTo struct {
	Email string `json:"email"`
}

type MailFrom struct {
	Email string `json:"email"`
}

type Personalization struct {
	To []MailTo `json:"to"`
}

type Mail struct {
	Personalizations []Personalization `json:"personalizations"`
	From             MailFrom          `json:"from"`
	Subject          string            `json:"subject"`
	Content          []MailContent     `json:"content"`
}

func sendNotification(apiKey string, email string, subject string, content string) error {
	mail := Mail{
		Personalizations: []Personalization{{
			To: []MailTo{{Email: email}},
		}},
		From:    MailFrom{Email: email},
		Subject: subject,
		Content: []MailContent{{
			Type:  "text/plain",
			Value: content,
		}},
	}
	data, err := json.Marshal(&mail)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", sendgridMailSendAPI, bytes.NewReader(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 202 {
		return fmt.Errorf("failed to send notification email: API responded with status code %d", response.StatusCode)
	}

	return nil
}

func downloadEpisode(dataDir string, episodeNumber int) error {
	magnet, err := getMagnet(episodeNumber)
	if err != nil {
		return fmt.Errorf("failed to get magnet: %w", err)
	}

	config := torrent.NewDefaultClientConfig()
	config.DataDir = dataDir
	cl, err := torrent.NewClient(config)
	if err != nil {
		return err
	}
	defer cl.Close()

	tfile, err := cl.AddMagnet(magnet)
	if err != nil {
		return fmt.Errorf("failed to add magnet: %w", err)
	}

	log.Println("getting an info for the torrent")
	<-tfile.GotInfo()
	log.Println("an info for the torrent has become available")

	tfile.DownloadAll()

	log.Println("waiting torrent client to complete all downloads")
	if cl.WaitAll() {
		log.Println("all torrents are completely downloaded")
	} else {
		return fmt.Errorf("torrent client has stopped before all torrents have been downloaded")
	}

	return nil
}

func latestDownloadedEpisode(dataDir string) (int, error) {
	file := path.Join(dataDir, ".latest_downloaded_episode")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		//log.Fatalf("Could not read latest downloaded episode: %v", err)
		return 0, err
	}
	episodeNumber, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, nil
	}
	return episodeNumber, err
}

func updateLatestDownloadedEpisode(dataDir string, episodeNumber int) error {
	file := path.Join(dataDir, ".latest_downloaded_episode")
	err := ioutil.WriteFile(file, []byte(strconv.Itoa(episodeNumber)), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getNextRun(now time.Time, interval time.Duration) time.Time {
	// TODO 	nextRun := time.Date(up.Year(), up.Month(), up.Day(), up.Hour(), 0, 0, 0, up.Location())
	return time.Now().Add(interval).Round(interval)
}

func main() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Fatalf("Environ variable SENDGRID_API_KEY must be set")
	}

	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		log.Fatalf("Environ variable DATA_DIR must be set")
	}

	scrapeIntervalRaw := os.Getenv("SCRAPE_INTERVAL")
	if scrapeIntervalRaw == "" {
		scrapeIntervalRaw = "1h"
	}

	notificationEmail := os.Getenv("NOTIFICATION_EMAIL")

	scrapeInterval, err := time.ParseDuration(scrapeIntervalRaw)
	if err != nil {
		log.Fatalf("Eviron variable SCRAPE_INTERVAL must be of type time.Duration")
	}

	for {
		nextRun := getNextRun(time.Now(), scrapeInterval)
		log.Printf("[INFO] Next try at %v", nextRun)
		<-time.After(nextRun.Sub(time.Now()))

		episodeNumber, err := latestDownloadedEpisode(dataDir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				episodeNumber = 889
			} else {
				log.Printf("[ERROR] Could not read latest downloaded episode number: %v", err)
				continue
			}
		}

		// try to download next episode
		episodeNumber += 1
		log.Printf("[INFO] Download episode %v", episodeNumber)

		if err := downloadEpisode(dataDir, episodeNumber); err != nil {
			if errors.Is(err, ErrEpisodeNotFound) {
				log.Printf("[INFO] Episode %v not found", episodeNumber)
			} else {
				log.Printf("[ERROR] Failed to download episode %v: %v", episodeNumber, err)
			}
			continue
		}
		if err := updateLatestDownloadedEpisode(dataDir, episodeNumber); err != nil {
			log.Printf("[ERROR] Failed to update latest episode %v: %v", episodeNumber, err)
		}

		subject := fmt.Sprintf("[Horrible Crawler] New episode of One Piece %v available for watching", episodeNumber)
		content := fmt.Sprintf("One Piece episode %v successfully downloaded", episodeNumber)

		if notificationEmail != "" {
			if err := sendNotification(apiKey, notificationEmail, subject, content); err != nil {
				fmt.Printf("[ERROR] failed to send notification: %v\n", err)
				continue
			}
		}
	}
}
