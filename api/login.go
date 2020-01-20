package api

import (
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//"golang.org/x/oauth2/jwt"
	"gotest/db"
)

func Login(c *gin.Context){
	//var user db.User
	var data db.User
	c.BindJSON(&data)
	msg := db.Checker(&data)
	//user.Username = data.Username
	//user.Passwd = data.Passwd
	c.JSON(200,msg)
}

