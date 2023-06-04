package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
	"github.com/phanlop12321/golang/handler"
	"github.com/phanlop12321/golang/util"
)

func RequireUser(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)
		min := len("Bearer ")
		if len(header) <= min {
			util.Error(c, http.StatusUnauthorized, errors.New("token is require"))
			return
		}
		token := header[min:]
		claims, err := handler.VerifyToken(token)
		if err != nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		user, err := db.GetUserByID(claims.UserID)
		if err != nil || user == nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		handler.SetUser(c, user)
	}
}
