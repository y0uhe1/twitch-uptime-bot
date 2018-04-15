package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetStreams ...
func GetStreams(userLogin string) Streams {
	values := url.Values{}
	values.Add("user_login", userLogin)

	url := "https://api.twitch.tv/helix/streams?" + values.Encode()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Client-ID", config.Twitch.ClientID)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err.Error())
		return Streams{}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err.Error())
		return Streams{}
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s", body)
		return Streams{}
	}

	var streams Streams
	if err := json.Unmarshal(body, &streams); err != nil {
		log.Println(err.Error())
		return Streams{}
	}

	defer resp.Body.Close()

	return streams
}

// GetTopClips TBD
func GetTopClips(userLogin string) Clips {
	values := url.Values{}
	values.Add("channel", userLogin)
	values.Add("period", "week")
	values.Add("limit", "1")

	url := "https://api.twitch.tv/kraken/clips/top?" + values.Encode()
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Client-ID", config.Twitch.ClientID)
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err.Error())
		return Clips{}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err.Error())
		return Clips{}
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s", body)
		return Clips{}
	}

	var clips Clips
	if err := json.Unmarshal(body, &clips); err != nil {
		log.Println(err.Error())
		return Clips{}
	}

	defer resp.Body.Close()

	return clips
}
