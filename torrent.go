package main

import (
	"fmt"
	"log"

	"github.com/anacrolix/torrent"
	"github.com/antchfx/htmlquery"
)

var ErrEpisodeNotFound = fmt.Errorf("episode not found")

func ShowURL(showID ShowID) string {
	return fmt.Sprintf(`https://horriblesubs.info/api.php?method=getshows&type=show&showid=%v`, showID)
}

func MagnetLink(showID ShowID, episodeNumber int) (string, error) {
	showURL := ShowURL(showID)
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

func DownloadEpisode(dataDir string, showID ShowID, episodeNumber int) error {
	magnet, err := MagnetLink(showID, episodeNumber)
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

	log.Println("[INFO] Getting an info for the torrent")
	<-tfile.GotInfo()
	log.Println("[INFO] An info for the torrent has become available")

	tfile.DownloadAll()

	log.Println("[INFO] Waiting torrent client to complete all downloads")
	if cl.WaitAll() {
		log.Println("[INFO] All torrents are completely downloaded")
	} else {
		return fmt.Errorf("torrent client has stopped before all torrents have been downloaded")
	}

	return nil
}
