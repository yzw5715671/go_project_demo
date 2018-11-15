package filters

import (
	"github.com/gin-gonic/gin"
	"golibs/cst"
)

func XxxBefore() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth
		err := xxxAuth(c)
		if err == nil {
			c.Next()
			return
		}

		// end
		sendResponse(c)
	}
}

func XxxAfter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// end
		sendResponse(c)
	}
}

func xxxAuth(c *gin.Context) (err error) {
	return
}

func sendResponse(c *gin.Context, ) {
	ret, ok := c.Get(cst.CTX_KEY_RET)
	if ok == false {
		ret = nil
	}

	c.JSON(200, ret)
	c.Writer.Flush()
	c.Abort()
}
