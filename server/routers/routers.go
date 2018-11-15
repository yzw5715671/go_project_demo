package routers

import (
"fmt"
"github.com/gin-gonic/gin"
	"go-demo-project/server/actions/aaa"
	"go-demo-project/server/actions/bbb"
	"go-demo-project/server/filters"
)

func Init(g *gin.Engine) (err error) {
	if g == nil {
		err = fmt.Errorf("nil gin engine")
		return
	}

	// root group
	rg := g.Group("xxx")
	rg.Use(filters.XxxBefore(), filters.XxxAfter())

	// aaa
	aaaG := rg.Group("aaa")
	{
		aaaG.GET("/get", aaa.Get)
	}

	// bbb
	bbbG := rg.Group("bbb")
	{
		bbbG.POST("/set", bbb.Set)
	}

	return
}

