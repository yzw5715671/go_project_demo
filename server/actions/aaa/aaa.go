package aaa

import (
	"github.com/gin-gonic/gin"
	"golibs/cst"
	"github.com/spf13/cast"
	"golibs/utils"
	"go-demo-project/lib/aaa"
)

func Get(c *gin.Context) {
	aaa := cast.ToInt(c.DefaultQuery("aaa", ""))

	// 业务逻辑
	ret, err := libaaa.Get(aaa)
	if err != nil {
		utils.SetRsp(c, cst.ERR_SRV_INTERNAL, err.Error(), err, nil)
	}

	//
	utils.SetRsp(c, cst.ERR_OK, "", nil, ret)
	return
}
