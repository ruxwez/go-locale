package gin

import (
	"github.com/gin-gonic/gin"
	quiklocale "github.com/quikcode/golocale"
)

// Middleware
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var locale string = quiklocale.Locale

		header := c.GetHeader(quiklocale.HeaderKey)
		if header != "" {
			locale = header
		}

		c.Set("locale", locale)

		c.Next()
	}
}

func GetMessage(c *gin.Context, code string) string {
	locale := c.MustGet("locale").(string)
	return quiklocale.GetMessage(locale, code)
}
