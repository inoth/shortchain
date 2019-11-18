package controller

import (
	"fmt"
	"net/http"
	ex "shortchain/controller/exception"
	"shortchain/model"
	"shortchain/model/reqmodel"
	"shortchain/services"
	"shortchain/util"
	"shortchain/util/config"
	"shortchain/util/result"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateShortChain(c *gin.Context) {
	req := &reqmodel.ChainAdd{}
	if err := c.BindJSON(req); err != nil {
		panic(ex.ParamErr("param error"))
	}

	mch := &model.MethodInfo{}
	if err := services.FindOne(bson.M{"appid": req.Appid, "status": 1}, mch); err != nil {
		panic(ex.DataNil("appid not found"))
	}

	chain := &model.ChainInfo{
		ShortId:    util.GuuidWithLens(8),
		Appid:      mch.Appid,
		LongChain:  req.LongChain,
		Status:     1,
		CreateTime: time.Now()}
	if err := services.Create(chain); err != nil {
		c.JSON(http.StatusOK, result.ResultErr("生成失败"))
	}
	c.JSON(http.StatusOK, result.ResultOK("ok", gin.H{"url": config.Instance().Domain + chain.ShortId}))
}

func RedirectTo(c *gin.Context) {
	shortid := c.Param("shortid")
	ip := c.GetHeader("X_REAL_IP")

	chain := &model.ChainInfo{}
	if err := services.FindOne(bson.M{"shortid": shortid, "status": 1}, chain); err != nil {
		// 404
		c.JSON(http.StatusNotFound, result.ResultNoData(result.NOTFOUND, "无效的短链"))
	}

	data := &model.UseRecord{
		Appid:      chain.Appid,
		ShortId:    shortid,
		LongChain:  chain.LongChain,
		HostIP:     ip,
		CreateTime: time.Now()}
	if err := services.Create(data); err != nil {
		fmt.Printf("保存跳转记录失败： %v", err)
	}
	c.Redirect(http.StatusMovedPermanently, chain.LongChain)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"code": 200})
}
