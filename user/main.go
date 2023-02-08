package main

import (
	api "github.com/joker-star-l/dousheng_idls/user/kitex_gen/api/user"
	"log"
)

func main() {
	svr := api.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
