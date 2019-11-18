package reqmodel

type Registered struct {
	Verify
	Account string `binding:"required"`
	Name    string `binding:"required"`
}
