package userapi

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// authorized is an interceptor
func (r *Controller) authorized() gin.HandlerFunc {

	return func(c *gin.Context) {

		bearToken := c.GetHeader("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		tk, err := r.UserToken.VerifyToken(strArr[1])
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if _, ok := tk.Claims.(jwt.Claims); !ok && !tk.Valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

	}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
