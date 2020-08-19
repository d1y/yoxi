package main

import (
	"fmt"

	"github.com/d1y/yoxi/freeport"
	"github.com/d1y/yoxi/server"
	"github.com/d1y/yoxi/utils"
	"github.com/d1y/yoxi/x"
)

func createWebViewURL(x int) string {
	return fmt.Sprintf("http://localhost:%v?target=desktop", x)
}

func main() {
	var free = freeport.GetPort()
	var url = createWebViewURL(free)
	var curr = utils.Curr()
	fmt.Println("curr", curr)
	go func() {
		fmt.Println("run web server..")
		server.CreateServer(free)
	}()
	x.CreateWebview(url)
}
