package main

import (
	"UserMockGo/config"
	"UserMockGo/web"
)

func main() {
	conf := config.ReadConfig()
	web.Init(conf)
}
