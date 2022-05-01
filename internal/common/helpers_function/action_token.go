package helpers_function

import (
	errorsCustom "News24/internal/app/control_users"
	"News24/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"

	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GetTokenByUser(user *models.User) (string, error) {

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	signKey := os.Getenv("SIGN_KEY")

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signKey))

}

func extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyToken(c *gin.Context) (err error) {

	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGN_KEY")), nil
	})
	if err != nil {
		return errorsCustom.InvalidAccessToken
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errorsCustom.TokenExpired
}

func GetUserByToken(c *gin.Context) (user *models.User, err error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGN_KEY")), nil
	})
	if err != nil {
		return user, errorsCustom.InvalidAccessToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		users, ok := claims["user"]
		if !ok {
			return user, errorsCustom.InvalidAccessToken
		}

		jsonUser, err := json.Marshal(users)
		if err != nil {
			return user, errorsCustom.InvalidAccessToken
		}
		user = &models.User{}

		json.Unmarshal(jsonUser, &user)

		return user, nil
	}
	return user, errorsCustom.TokenExpired
}
