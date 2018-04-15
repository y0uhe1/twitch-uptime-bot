package main

// Config ...
type Config struct {
	Server  ServerConfig  `toml:"server"`
	Twitch  TwitchConfig  `toml:"twitch"`
	Twitter TwitterConfig `toml:"twitter"`
}

// ServerConfig ...
type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// TwitchConfig ...
type TwitchConfig struct {
	ClientID   string   `toml:"client_id"`
	UserLogins []string `toml:"user_logins"`
}

// TwitterConfig ...
type TwitterConfig struct {
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerSecret    string `toml:"consumer_secret"`
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
}
