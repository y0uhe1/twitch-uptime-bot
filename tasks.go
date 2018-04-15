package main

import (
	"fmt"
	"log"
	"time"
)

// CheckStreamAndTweet ...
func CheckStreamAndTweet() {
	log.Print("処理開始")

	userLogins := config.Twitch.UserLogins

	for _, userLogin := range userLogins {

		log.Printf("Check %s's stream ... ", userLogin)

		checktimes := 0
		var streams Streams

		// Twitchはゴールデンタイムはつながらないことがあるので5回リトライする
		for checktimes < 5 {
			streams = GetStreams(userLogin)
			if len(streams.Data) == 0 {
				checktimes++

				// 一秒待ってからTwitchリトライ
				time.Sleep(1 * time.Second)
			} else {
				// streamsにデータが入っていればループは抜ける
				checktimes = 5
			}
		}

		for _, data := range streams.Data {
			log.Printf("%s is online.", userLogin)

			uptime := time.Since(data.StartedAt)
			hour := int(uptime.Hours())
			sec := int(uptime.Seconds())
			rem := sec % 3600

			log.Printf("current hour: %d, sec: %d, rem: %d.\n", hour, sec, rem)

			if rem < 300 {

				log.Printf("rem is less than 300 sec. Set up for tweet.")

				var prefix string
				if hour == 0 {
					prefix = "配信開始"
				} else {
					prefix = fmt.Sprintf("%d時間経過", hour)
				}

				status := fmt.Sprintf("Twitchで配信中！\n\n【%s】%s\nhttps://www.twitch.tv/%s", prefix, data.Title, userLogin)

				log.Printf("%s is streaming for %d hour.\n", userLogin, hour)
				Tweet(status, data.ThumbnailURL)
			}
		}
		log.Printf("%s is offline. Skip tweet.", userLogin)

	}
	log.Print("処理終了")
}

// CheckClipAndTweet TBD
func CheckClipAndTweet(userLogin string) {
	clips := GetTopClips(userLogin)

	for _, clip := range clips.Clips {
		now := time.Now()
		today := now.Format("2006/01/02")
		aweekbefore := now.AddDate(0, 0, -7).Format("2006/01/02")
		status := fmt.Sprintf("【今週(%s~%s)の人気クリップ by %s】\n\n%s\n%s\n\n▶チャンネルはこちら https://www.twitch.tv/%s", aweekbefore, today, userLogin, clip.Title, clip.URL, userLogin)
		Tweet(status, clip.Thumbnails.Medium)
	}

}
