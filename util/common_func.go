package util

import (
	"crypto/md5"
	"encoding/hex"
	ex "shortchain/controller/exception"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GuuidWithLens(length int) string {
	buf := hex.EncodeToString(uuid.NewV4().Bytes()[:])
	return buf[:length]
}

func GetTimespan() int64 {
	return time.Now().Unix()
}

func Md5(origin string) string {
	buf := md5.Sum([]byte(origin))
	return hex.EncodeToString(buf[:])
}

func ToString(o interface{}) (string, error) {
	switch o.(type) {
	case int:
		return strconv.Itoa(o.(int)), nil
	case int64:
		return strconv.FormatInt(o.(int64), 10), nil
	case float64:
		return strconv.FormatFloat(o.(float64), 'E', -1, 64), nil
	case string:
		return o.(string), nil
	default:
		return "", &ex.SystemErr{Msg: "can't change this type"}
	}
}
