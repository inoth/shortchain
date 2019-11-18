package reqmodel

type Verify struct {
	Timespan int64  `binding:"required"`
	Noncestr string `binding:"required"`
	Sign     string `binding:"required"`
}
