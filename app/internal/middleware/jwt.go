package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"ilanver/internal/helpers"
	"ilanver/internal/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// APIError struct
type APIError struct {
	Code    int
	Message string
}

// RespondWithError function handler error
func RespondWithError(code int, message string, c *gin.Context) {
	c.JSON(code, &APIError{Code: code, Message: message})
	c.Abort()
}

// JWTAuthMiddleware middleware function implementation
func JWTAuthMiddleware(encoded bool, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)

		// Token base validation
		if auth == "" {
			RespondWithError(401, "API token required", c)
			return
		}

		// Token base validation
		if strings.ToLower(parts[0]) != "bearer" {
			RespondWithError(401, "Authorization header must start with Bearer", c)
			return
		} else if len(parts) == 1 {
			RespondWithError(401, "Token not found", c)
			return
		} else if len(parts) > 2 {
			RespondWithError(401, "Authorization header must be Bearer and token", c)
			return
		}

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing me thod: %v", token.Header["alg"])
			}

			mySecertKey := secret

			if encoded {
				var replacer = strings.NewReplacer("_", "/", "-", "+")
				mySecertKey = replacer.Replace(mySecertKey)

				base64Decoded, _ := base64.StdEncoding.DecodeString(mySecertKey)
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return base64Decoded, nil
			}

			return []byte(mySecertKey), nil

		})

		claims := token.Claims.(jwt.MapClaims)
		tmpUser := claims["user"]
		tmp, _ := json.Marshal(tmpUser)
		user := model.User{}
		_ = json.Unmarshal(tmp, &user)

		helpers.AuthUser = user

		if token.Valid {
			c.Next()
		} else {
			RespondWithError(401, err.Error(), c)
			return
		}
	}
}
