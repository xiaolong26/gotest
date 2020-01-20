package router

import (
	"gotest/api"
	"github.com/gin-gonic/gin"
	"gotest/jwt"
	)
func SetupRouter(r *gin.Engine) *gin.Engine {
	//rr := r.Group("/")
	r.POST("/login",api.Login)
	authrization := r.Group("/",jwt.JWTAuth())
	{
		authrization.GET("user",api.Test)
	}
	return r
}