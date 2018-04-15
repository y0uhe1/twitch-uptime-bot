# twitch-uptime-bot

## 概要
チェックしたいTwitch配信者が
* 配信を開始したとき 
* 配信後1時間毎 

に配信情報をTwitterで呟くBOTです

### 準備しておくこと
* Twitchの開発者に登録しクライアントIDを入手しておく
* Twitterの開発者に登録し各種キーを入手しておく
* config.tml の設定が必要

### config.tml
```toml:config.tml
[server]
host = "localhost"
port = "8080"

[twitch]
client_id = "TWITCH_CLIENT_ID"
user_logins = ["USER_LOGIN"] #チェックしたいTwichのユーザー名の配列

[twitter]
consumer_key = "CONSUMER_KEY" 
consumer_secret = "CONSUMER_SECRET" 
access_token = "ACCESS_TOKEN" 
access_token_secret = "ACCESS_TOKEN_SECRET" 
```

### 利用方法
1. 任意のディレクトリにクローン
2. `go build`
3. ビルドされた実行ファイルを実行
4. http://localhost:8080/start でジョブスタート
5. http://localhost:8080/stop でジョブスタート

### 利用しているBOT
Tweetのイメージは利用しているBOTのつぶやきを参照

https://twitter.com/dtn_stream_bot

https://twitter.com/dtn_kr_stream

興味ある方は是非フォローお願いします