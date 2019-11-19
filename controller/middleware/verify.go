package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortchain/util"
	"shortchain/util/config"
	"shortchain/util/result"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequestVerify(c *gin.Context) {

	cp := c.Copy()
	ch := make(chan int, 1)
	go func() {
		defer close(ch)

		data, err := cp.GetRawData()
		if err != nil {
			fmt.Printf("get data error: %v", err)
			ch <- result.FAILED
			return
		}

		v := make(map[string]interface{})
		err = json.Unmarshal(data, &v)
		if err != nil {
			fmt.Printf("convert error: %v", err)
			ch <- result.FAILED
			return
		}
		// 重复请求过滤 redis
		// 时间戳校验
		if v["timespan"] == nil || v["sign"] == nil || v["noncestr"] == nil {
			ch <- result.PARAMERR
			return
		}

		ts, ok := v["timespan"].(float64)
		if !ok {
			ch <- result.PARAMERR
			return
		}
		if (util.GetTimespan() - int64(ts)) > 180 {
			ch <- result.VERIFYERR
			return
		}

		// 签名校验
		if !checkSign(v) {
			ch <- result.VERIFYERR
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		ch <- result.SUCCESS
	}()

	switch <-ch {
	case result.SUCCESS:
		c.Next()
	case result.FAILED:
		c.Abort()
		c.JSON(http.StatusInternalServerError, result.ResultNoData(result.FAILED, "系统繁忙"))
	case result.VERIFYERR:
		c.Abort()
		c.JSON(http.StatusOK, result.ResultNoData(result.VERIFYERR, "拒绝访问"))
	case result.PARAMERR:
		c.Abort()
		c.JSON(http.StatusOK, result.ResultNoData(result.VERIFYERR, "参数异常"))
	}
}

func checkSign(m map[string]interface{}) bool {

	sign := m["sign"].(string)

	var keys []string
	for k := range m {
		if strings.EqualFold(k, "sign") {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	for _, k := range keys {
		buffer.WriteString(k)
		buffer.WriteString("=")
		if s, err := util.ToString(m[k]); err == nil {
			buffer.WriteString(s)
		}
		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(config.Instance().Token)

	md := util.Md5(buffer.String())
	return strings.EqualFold(md, sign)
}
