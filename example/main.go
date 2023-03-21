package main

import (
	server "github.com/gin-gonic/gin"
	"github.com/quikcode/quiklocale-go"
	mw "github.com/quikcode/quiklocale-go/gin"
)

func main() {
	r := server.New()
	quiklocale.Init("en", "Accept-Language", []string{
		"./locales/en.json",
		"./locales/es.json",
	})

	r.Use(mw.Middleware())

	r.GET("/test", func(c *server.Context) {
		c.JSON(200, mw.GetMessage(c, "test"))
	})

	// Listen and serve on 0.0.0.0:80
	r.Run(":80")
}
