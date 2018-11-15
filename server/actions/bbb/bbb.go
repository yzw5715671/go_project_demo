package bbb

import (
	"github.com/gin-gonic/gin"
	"go-demo-project/lib/bbb"
	"golibs/cst"
	"github.com/alecthomas/log4go"
	"golibs/utils"
)

func Set(c *gin.Context) {

	var req libbbb.SetInput
	err := utils.GetPostBody(c, &req)
	if err != nil {
		utils.SetRsp(c, cst.ERR_CLT_INVALID_PARAM, err.Error(), err, nil)
		return
	}
	log4go.Info("req=%v", req)

	//
	ret, err := libbbb.Set(req)
	if err != nil {
		utils.SetRsp(c, cst.ERR_SRV_DB, err.Error(), err, nil)
		return
	}

	utils.SetRsp(c, cst.ERR_OK, "成功", nil, ret)
	return
}
