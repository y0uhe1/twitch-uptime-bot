package main

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

// Tweet ...
func Tweet(status string, thumbnailURL string) {
	tmpfile, err := ioutil.TempFile("", "tmp")

	if err != nil {
		log.Println(err.Error())
		return
	}

	r := strings.NewReplacer("{width}", "1280", "{height}", "720")
	imgURL := r.Replace(thumbnailURL)
	resp, err := http.Get(imgURL)

	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	defer tmpfile.Close()

	io.Copy(tmpfile, resp.Body)

	filedata, _ := ioutil.ReadFile(tmpfile.Name())
	imgEnc := base64.StdEncoding.EncodeToString(filedata)

	anaconda.SetConsumerKey(config.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(config.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)

	media, _ := api.UploadMedia(imgEnc)
	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)
	result, err := api.PostTweet(status, v)

	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("Post tweet %s\n", result.Text)
}
