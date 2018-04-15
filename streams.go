package main

import "time"

// Streams ...
type Streams struct {
	Data []struct {
		ID           string        `json:"id"`
		UserID       string        `json:"user_id"`
		GameID       string        `json:"game_id"`
		CommunityIds []interface{} `json:"community_ids"`
		Type         string        `json:"type"`
		Title        string        `json:"title"`
		ViewerCount  int           `json:"viewer_count"`
		StartedAt    time.Time     `json:"started_at"`
		Language     string        `json:"language"`
		ThumbnailURL string        `json:"thumbnail_url"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// Clips ...
type Clips struct {
	Clips []struct {
		Slug        string `json:"slug"`
		TrackingID  string `json:"tracking_id"`
		URL         string `json:"url"`
		EmbedURL    string `json:"embed_url"`
		EmbedHTML   string `json:"embed_html"`
		Broadcaster struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
			ChannelURL  string `json:"channel_url"`
			Logo        string `json:"logo"`
		} `json:"broadcaster"`
		Curator struct {
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			DisplayName string      `json:"display_name"`
			ChannelURL  string      `json:"channel_url"`
			Logo        interface{} `json:"logo"`
		} `json:"curator"`
		Vod struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"vod"`
		Game       string    `json:"game"`
		Language   string    `json:"language"`
		Title      string    `json:"title"`
		Views      int       `json:"views"`
		Duration   float64   `json:"duration"`
		CreatedAt  time.Time `json:"created_at"`
		Thumbnails struct {
			Medium string `json:"medium"`
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
		} `json:"thumbnails"`
	} `json:"clips"`
	Cursor string `json:"_cursor"`
}
