package main

import (
	"os"

	"github.com/mazrean/todoList/common"
)

func main() {
	addr, ok := os.LookupEnv("PORT")
	if !ok {
		panic("PORT is not set")
	}

	sessionKey, ok := os.LookupEnv("SESSION_KEY")
	if !ok {
		panic("SESSION_KEY is not set")
	}

	sessionSecret, ok := os.LookupEnv("SESSION_SECRET")
	if !ok {
		panic("SESSION_SECRET is not set")
	}

	api, err := InjectAPI(&Config{
		Addr:          common.Addr(addr),
		SessionKey:    common.SessionKey(sessionKey),
		SessionSecret: common.SessionSecret(sessionSecret),
	})
	if err != nil {
		panic(err)
	}

	api.Start()
}
