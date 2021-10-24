package jwt

import (
	"blogapi/pkg/e"
	"blogapi/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized,
				unauthResponse(e.INVALID_PARAMS))
			// TODO Abort
			c.Abort()
			return
		}

		claim, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized,
				unauthResponse(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
			c.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			c.JSON(http.StatusUnauthorized,
				unauthResponse(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT))
			c.Abort()
			return
		}

		c.Next()
	}
}

func unauthResponse(code int) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	}
}
