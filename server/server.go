package server

import (
	"log"

	"github.com/d1y/yoxi/config"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/pkg/errors"
)

// CreateServer 创建服务
func CreateServer(port int) {

	var xPort = port

	if port == 0 {
		xPort = config.WebServerDefaultPort
	}

	app := fiber.New(&fiber.Settings{
		// Prefork: true,
	})

	// app.Use(func(c *fiber.Ctx) {
	// 	c.SendStatus(404)
	// })

	// test server
	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("ok")
	})

	app.Static("/", config.WebDistPath)
	app.Static("/results", config.WebAssetsPath)

	app.Use(middleware.Logger())

	e := app.Listen(xPort)

	if e != nil {
		err := errors.Wrap(e, "create web server is error")
		log.Fatalf("%+v\n", err)
	}

}
