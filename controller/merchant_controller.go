package controller

import (
	"net/http"
	ex "shortchain/controller/exception"
	"shortchain/model"
	"shortchain/model/reqmodel"
	"shortchain/services"
	"shortchain/util"
	"shortchain/util/result"
	"time"

	"github.com/gin-gonic/gin"
)

func Registered(c *gin.Context) {
	req := &reqmodel.Registered{}
	if err := c.BindJSON(req); err != nil {
		panic(ex.ParamErr("param error"))
	}

	mch := &model.MethodInfo{
		Appid:      "ino" + util.GuuidWithLens(16),
		Account:    req.Account,
		Name:       req.Name,
		Status:     1,
		CreateTime: time.Now()}
	if err := services.Create(mch); err != nil {
		c.JSON(http.StatusOK, result.ResultErr("注册失败"))
	}
	c.JSON(http.StatusOK, result.ResultOK("ok", gin.H{"appid": mch.Appid}))
}
