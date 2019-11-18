package middleware

import (
	"fmt"
	ex "shortchain/controller/exception"

	"shortchain/util/result"

	"github.com/gin-gonic/gin"
)

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 捕获错误
			switch e := err.(type) {
			case *ex.ParamErrException:
				fmt.Printf("参数错误 %v", err)
				c.JSON(200, result.ResultNoData(result.FAILED, e.Msg))
			case *ex.DataNilException:
				fmt.Printf("空数据 %v", err)
				c.JSON(200, result.ResultNoData(result.FAILED, e.Msg))
			case *ex.VerifyErrException:
				fmt.Printf("请求验证错误 %v", err)
				c.JSON(200, result.ResultNoData(result.FAILED, e.Msg))
			default:
				fmt.Printf("出现未知错误 %v", err)
				c.JSON(500, result.ResultNoData(result.FAILED, "unknown exception."))
			}
		}
	}()
	c.Next()
}
