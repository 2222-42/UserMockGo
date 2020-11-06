package config

import "os"

type NotifierConfig struct {
	FromAddress    string
	SendGridApiKey string
}

type Config struct {
	NotifierConfig *NotifierConfig
}

func ReadConfig() *Config {
	config := Config{NotifierConfig: &NotifierConfig{}}
	config.setNotifierConfig()

	return &config
}

func (conf *Config) setNotifierConfig() {
	conf.NotifierConfig.FromAddress = os.Getenv("FROM_ADDRESS")
	conf.NotifierConfig.SendGridApiKey = os.Getenv("SENDGRID_API_KEY")
}
