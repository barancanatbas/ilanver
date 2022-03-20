package helpers

import (
	"encoding/json"
	"fmt"
	"ilanver/internal/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var AuthUser model.User

// buna gerek yok gibi.
func Auth(c *gin.Context) model.User {
	auth := c.GetHeader("Authorization")
	parts := strings.Fields(auth)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing me thod: %v", token.Header["alg"])
		}

		mySecertKey := "secret"

		return []byte(mySecertKey), nil
	})

	claims := token.Claims.(jwt.MapClaims)
	tmpUser := claims["user"]
	tmp, _ := json.Marshal(tmpUser)
	user := model.User{}
	_ = json.Unmarshal(tmp, &user)

	return user
}
