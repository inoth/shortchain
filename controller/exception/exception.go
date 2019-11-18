package exception

type ParamErrException struct {
	Msg string
}

func (ex *ParamErrException) Error() string {
	return ex.Msg
}

func ParamErr(msg string) *ParamErrException {
	return &ParamErrException{Msg: msg}
}

type DataNilException struct {
	Msg string
}

func (ex *DataNilException) Error() string {
	return ex.Msg
}

func DataNil(msg string) *DataNilException {
	return &DataNilException{Msg: msg}
}

type VerifyErrException struct {
	Msg string
}

func (ex *VerifyErrException) Error() string {
	return ex.Msg
}

func VerifyErr(msg string) *VerifyErrException {
	return &VerifyErrException{Msg: msg}
}

type SystemErr struct {
	Msg string
}

func (ex *SystemErr) Error() string {
	return ex.Msg
}
