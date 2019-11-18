package model

import (
	"time"
)

type IEntity interface {
	ColName() string
}

type MethodInfo struct {
	Appid      string
	Account    string
	Name       string
	Status     uint
	CreateTime time.Time
}

func (u *MethodInfo) ColName() string {
	return "method_info"
}
