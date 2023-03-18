package main

import (
	"go-boilerplate-v3/domains/user/core-api"
	"go-boilerplate-v3/domains/user/created-log-consumer"
	"log"
	"os"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	var appName = os.Args[1]

	if appName == "user-core-api" {
		core_api.Init()
	} else if appName == "user-created-log-consumer" {
		created_log_consumer.Init()
	}
}
