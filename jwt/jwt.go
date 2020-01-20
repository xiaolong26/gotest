package jwt
import (
	"github.com/gin-gonic/gin"
	"log"

	//"github.com/dgrijalva/jwt-go"
	"net/http"
)

func JWTAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "无token,拒绝访问",
			})
			c.Abort()
			return
		}
		log.Print("get token :",token)

	}
}