package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	rssFeed := RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &rssFeed, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req.Header.Add("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return &rssFeed, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &rssFeed, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &rssFeed, err
	}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &rssFeed, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString((rssFeed.Channel.Description))

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}
