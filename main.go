package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.UserRoute(routes)
	router.Run(":8080")

}
