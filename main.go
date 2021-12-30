package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "./avatars/default.png")
	r.StaticFile("/favicon.ico", "./avatars/default.png")
	r.GET("/:id", getAvatar)

	r.Run()
}
