package main

import (
	"github.com/gin-gonic/gin"
	"gotest/router"
)

func main(){
	r := gin.Default()
	rr := router.SetupRouter(r)
	rr.Run()
}