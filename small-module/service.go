package main

import "github.com/kamioair/qf"

type service struct {
	qf.Service

	// 具体业务功能实现
	bll *bll
}

// Reg 注册需要执行的方法
func (b *service) Reg(reg *qf.Reg) {
	reg.OnInit = b.onInit
}

func (b *service) onInit() {
	b.bll = newBll(b.CreateCron())
}
