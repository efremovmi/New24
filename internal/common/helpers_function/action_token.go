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
	if token == "" {
		tokenSplit := strings.Split(c.GetHeader("Cookie"), "token=")
		if len(tokenSplit) == 2 {
			return strings.Split(tokenSplit[1], ";")[0]
		}
		token = ""
	}
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
			return nil, errorsCustom.InvalidAccessToken
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

func getUserByToken(c *gin.Context) (user *models.User, err error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorsCustom.InvalidAccessToken
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

func IsModerator(c *gin.Context) error {
	user, err := getUserByToken(c)
	if err != nil {
		return err
	}

	role, err := strconv.Atoi(fmt.Sprintf("%v", user.Role))
	if err != nil {
		return errorsCustom.Forbidden
	}

	if role < 1 {
		return errorsCustom.Forbidden
	}

	return nil
}

func IsAdmin(c *gin.Context) (userAdmin *models.User, err error) {
	userAdmin, err = getUserByToken(c)
	if err != nil {
		return nil, err
	}

	role, err := strconv.Atoi(fmt.Sprintf("%v", userAdmin.Role))

	if err != nil {
		return nil, errorsCustom.Forbidden
	}

	if role < 2 {
		return nil, errorsCustom.Forbidden
	}

	return userAdmin, nil
}
