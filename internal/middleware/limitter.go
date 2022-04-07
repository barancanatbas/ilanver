package middleware

import (
	"fmt"
	"ilanver/internal/cache"

	"github.com/gin-gonic/gin"
)

// gelen istekleri sınırlandırmak için burada bir middleware oluşturuyoruz.
// bunu redis ile yapıyoruz çünkü redis çok hızlı :)
func Limitter() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		ip := c.ClientIP()
		if val, exists := cache.GetFromCacheInt(ip); exists {
			fmt.Println("ip: ", ip, " val: ", val)
			if val >= 200 {

				c.JSON(429, &APIError{Code: 429, Message: "Too many requests"})
				c.AbortWithStatus(429)
				return
			}
			// buradaki süreleri arttırabiliriz.
			cache.SetFromCache(ip, val+1, 60)
		} else {
			cache.SetFromCache(ip, 1, 60)
		}

		c.Next()
	}
}
