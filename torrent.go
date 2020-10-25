package main

import (
	"fmt"
	"log"

	"golang.org/x/net/html"

	"github.com/anacrolix/torrent"
	"github.com/antchfx/htmlquery"
)

var ErrEpisodeNotFound = fmt.Errorf("episode not found")

func ShowURL(episodeNumber int) string {
	return fmt.Sprintf(`https://nyaa.iss.one/?f=0&c=0_0&q=%%5BSubsPlease%%5D+One+Piece+-+%v+%%281080p%%29`, episodeNumber)
}

func MagnetLink(episodeNumber int) (string, error) {
	showURL := ShowURL(episodeNumber)
	node, err := htmlquery.LoadURL(showURL)
	if err != nil {
		return "", fmt.Errorf("failed to load url %q: %w", showURL, err)
	}

	return ExtractMagnetLink(node)
}

func ExtractMagnetLink(node *html.Node) (string, error) {
	tbody := htmlquery.FindOne(node, "//tbody")
	if tbody == nil {
		return "", fmt.Errorf("tbody not found: %w", ErrEpisodeNotFound)
	}

	td := htmlquery.FindOne(tbody, `//td[@class="text-center"]`)
	if td == nil {
		return "", fmt.Errorf("td not found")
	}

	links := htmlquery.Find(td, "//a")
	if len(links) == 0 {
		return "", fmt.Errorf("links not found")
	}

	href := htmlquery.SelectAttr(links[1], "href")
	if href == "" {
		return "", fmt.Errorf("href not found")
	}

	return href, nil
}

func DownloadEpisode(dataDir string, episodeNumber int) error {
	magnet, err := MagnetLink(episodeNumber)
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
