package main

import "github.com/FelizYK/gomall/app/email/mq"

func main() {
	mq.Init()
	defer mq.Close()

	select {}
}
