package main

import (
	"log"

	"github.com/d1y/yoxi/server"
)

func main() {
	log.Println("start web server")
	server.CreateServer(2333)
}
