package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var imgExt = []string{"png", "jpg", "jpeg", "gif"}

func file_exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getAvatar(c *gin.Context) {
	av := c.Param("id")
	for _, ext := range imgExt {
		file := fmt.Sprintf("./avatars/%s.%s", av, ext)
		if file_exists(file) {
			c.File(file)
			return
		}
	}

	file := "./avatars/default.png"
	if file_exits(file) {
	    c.File(file)
	    return
	}
	
	c.String(http.StatusNotFound, "Avatar not found")
}
