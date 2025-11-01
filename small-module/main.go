package main

import "github.com/kamioair/qf"

const (
	version = "V1.00.251101B01"
	module  = "SmallModule"
	desc    = "小型模块实例"
)

func main() {
	// 创建模块
	module := qf.NewModule(module, desc, version, &service{}, &config)
	// 启动
	module.Run()
}
