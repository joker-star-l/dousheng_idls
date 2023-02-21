package main

import (
	api "github.com/joker-star-l/dousheng_idls/message/kitex_gen/api/message"
	"log"
)

func main() {
	svr := api.NewServer(new(MessageImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
