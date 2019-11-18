package model

import (
	"time"
)

type UseRecord struct {
	Appid      string
	ShortId    string
	LongChain  string
	HostIP     string
	CreateTime time.Time
}

func (u *UseRecord) ColName() string {
	return "use_record"
}
