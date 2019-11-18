package model

import (
	"time"
)

type ChainInfo struct {
	ShortId    string
	Appid      string
	LongChain  string
	Status     uint
	CreateTime time.Time
}

func (u *ChainInfo) ColName() string {
	return "chain_info"
}
