package main

import (
	"fmt"
	"time"

	"github.com/kamioair/qf"
	"github.com/kamioair/utils/qos"
)

// mqttClientStatus MQTT客户端状态跟踪
type mqttClientStatus struct {
	failureCount  int       // 连续失败次数
	lastAlertTime time.Time // 最后一次报警时间
}

type bll struct {
	cron         qf.ICron
	clientStatus map[string]*mqttClientStatus // 客户端状态映射
}

func newBll(cron qf.ICron) *bll {
	b := &bll{
		cron:         cron,
		clientStatus: make(map[string]*mqttClientStatus),
	}
	// 定时检测mqtt客户端是否存在
	b.cron.Add("0 0/1 * * * ?", b.checkMqttClient)
	return b
}

func (b *bll) checkMqttClient() {
	for _, name := range config.CheckMqttClient {
		// 获取客户端状态，如果不存在则初始化
		status, exists := b.clientStatus[name]
		if !exists {
			status = &mqttClientStatus{
				failureCount:  0,
				lastAlertTime: time.Time{}, // 零值表示从未报警
			}
			b.clientStatus[name] = status
		}

		// 检测客户端是否存在
		exist := qos.CheckEmqxClientExist(
			config.CheckEmqx.Url,
			config.CheckEmqx.ApiKey,
			config.CheckEmqx.SecretKey,
			name)

		if exist == false {
			// 客户端不存在，增加失败计数
			status.failureCount++

			// 检查是否需要报警：连续失败超过配置阈值且距离上次报警超过冷却时间
			if status.failureCount > config.AlertConfig.MaxFailureCount && b.shouldAlert(status) {
				b.sendAlert(name)
				status.lastAlertTime = time.Now()
			}
		} else {
			// 客户端存在，重置失败计数
			status.failureCount = 0
		}
	}
}

// shouldAlert 检查是否应该报警（距离上次报警超过冷却时间）
func (b *bll) shouldAlert(status *mqttClientStatus) bool {
	// 如果从未报警过，应该报警
	if status.lastAlertTime.IsZero() {
		return true
	}

	// 检查距离上次报警是否超过配置的冷却时间
	cooldownDuration := time.Duration(config.AlertConfig.CooldownMinutes) * time.Minute
	return time.Since(status.lastAlertTime) > cooldownDuration
}

// sendAlert 发送报警
func (b *bll) sendAlert(clientName string) {
	fmt.Println("报警，客户端:", clientName)
}
