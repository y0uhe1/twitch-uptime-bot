package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/robfig/cron"
)

var c = cron.New()

var config Config

func init() {
	log.SetPrefix("[INFO]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 設定を読み込む
	log.Print("load config")
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	// cronにジョブを追加
	err = c.AddFunc("0 */5 * * * *", func() { CheckStreamAndTweet() })

	if err != nil {
		log.Print(err.Error())
	}

}

// StartCron is to start cron jobs
func StartCron(w http.ResponseWriter, r *http.Request) {
	c.Start()
	log.Print("Jobs start.")
	fmt.Fprintf(w, "Jobs start.")
}

// StopCron is to stop cron jobs
func StopCron(w http.ResponseWriter, r *http.Request) {
	c.Stop()
	log.Print("Jobs stop.")
	fmt.Fprintf(w, "Jobs stop.")
}

// Test ...
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, My Jobs!!")
}

func main() {
	http.HandleFunc("/", Test)
	http.HandleFunc("/start", StartCron)
	http.HandleFunc("/stop", StopCron)

	//http.ListenAndServeで待ち受けるportを指定
	err := http.ListenAndServe(":8080", nil)

	// エラー処理
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
