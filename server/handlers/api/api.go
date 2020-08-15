package api

import (
	"fmt"

	"github.com/d1y/yoxi/assets"
	"github.com/d1y/yoxi/audio"
	"github.com/gofiber/fiber"
)

// Play 播放
func Play(c *fiber.Ctx) {
	run := audio.Play(assets.Mp3)
	var id = run.Cmd.Process.Pid
	go func() {
		fmt.Println("该id为: ", id)
	}()
	c.SendString(fmt.Sprintf("%v", id))
}
