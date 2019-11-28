package result

const (
	SUCCESS   = 200
	PARAMERR  = 400
	VERIFYERR = 403
	NOTFOUND  = 404
	FAILED    = 500
)

type ApiResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ApiResultWithData struct {
	ApiResult
	Data interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}) *ApiResultWithData {
	return &ApiResultWithData{
		ApiResult: ApiResult{
			Code: code,
			Msg:  msg},
		Data: data}
}

func ResultNoData(code int, msg string) *ApiResult {
	return &ApiResult{Code: code, Msg: msg}
}

func ResultErr(msg string) *ApiResult {
	return &ApiResult{Code: FAILED, Msg: msg}
}

func ResultOKNoData(msg string) *ApiResult {
	return &ApiResult{Code: SUCCESS, Msg: msg}
}

func ResultOK(msg string, data interface{}) *ApiResultWithData {
	return &ApiResultWithData{
		ApiResult: ApiResult{
			Code: SUCCESS,
			Msg:  msg},
		Data: data}
}
