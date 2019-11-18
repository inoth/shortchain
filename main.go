package main

import (
	"fmt"
	"os"
	"shortchain/controller/router"
	"shortchain/db"
	"shortchain/util/config"
)

func init() {

	// 初始化配置
	if err := config.Instance().Init(); err != nil {
		fmt.Printf("init config is error: %v", err)
		os.Exit(1)
	}

	// 初始化数据库
	if err := db.Instance().Init("shortchain"); err != nil {
		fmt.Printf("init db is error: %v", err)
		os.Exit(1)
	}

	// 初始化路由
	router.Init()
}

func main() {
}
