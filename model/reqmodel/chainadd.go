package reqmodel

type ChainAdd struct {
	Verify
	Appid     string `binding:"required"`
	LongChain string `binding:"required"`
}
