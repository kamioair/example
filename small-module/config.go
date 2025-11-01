package main

import "github.com/kamioair/qf"

// config 配置定义
var config = struct {
	qf.Config

	// 自己模块的配置

	CheckMqttClient []string `comment:"需要检测的MQTT客户端名称列表"`
	CheckEmqx       struct {
		Url       string
		ApiKey    string
		SecretKey string
	} `comment:"EMQX配置\n Url:EMQX地址\n ApiKey/Secret:EMQX API密钥"`
	AlertConfig struct {
		MaxFailureCount int
		CooldownMinutes int
	} `comment:"报警配置\n MaxFailureCount:连续失败次数阈值，超过此值才报警\n CooldownMinutes:报警冷却时间（分钟），在此时间内不会重复报警"`
}{
	// 默认值
	CheckMqttClient: make([]string, 0),
	CheckEmqx: struct {
		Url       string
		ApiKey    string
		SecretKey string
	}{
		Url:       "127.0.0.1:18083",
		ApiKey:    "",
		SecretKey: "",
	},
	AlertConfig: struct {
		MaxFailureCount int
		CooldownMinutes int
	}{
		MaxFailureCount: 2,  // 连续失败超过2次开始报警
		CooldownMinutes: 60, // 1小时冷却时间
	},
}
