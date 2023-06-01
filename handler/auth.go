package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
	"github.com/phanlop12321/golang/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost      = 12
	secretKey = "SuperSecret"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(RegisterReq)
		if err := c.BindJSON(req); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), cost)
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		user := model.User{
			Username: req.Username,
			Password: string(hash),
		}
		if err := db.CreateUser(&user); err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		token, err := GenerateToken(user.ID)
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	payload := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "course-api",
		},
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return claim.SignedString([]byte(secretKey))
}

func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
