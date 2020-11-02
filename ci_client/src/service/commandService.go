package service

import (
	"ci_client/command"
	"ci_client/common"
	"sync"
)

type commandService struct {
	cmd map[int]*command.Command
}

var instance *commandService
var once sync.Once

// 获取单例
func GetInstance() *commandService {
	once.Do(func() {
		instance = &commandService{}
	})
	return instance
}

func (service *commandService) GetInfo() common.ResponseEntity {
	doc := common.New()
	doc.Start()
	doc.AddStr("Info", "Info")
	doc.End()
	return common.ResponseEntity{State: common.SUCCESS, Data: doc.String()}
}

func (service *commandService) Start(evnConf string, buildConf string, id int) common.ResponseEntity {
	// 创建
	cmd := command.New(evnConf, buildConf)
	// 开始
	cmd.Start()
	// 保存
	service.cmd = make(map[int]*command.Command)
	service.cmd[id] = cmd
	return common.OK
}

func (service *commandService) Stop(id int) common.ResponseEntity {
	if v, ok := service.cmd[id]; ok {
		// 存在
		v.Kill()
		return common.OK
	}
	return common.Failed
}

func (service *commandService) Download() common.ResponseEntity {
	doc := common.New()
	doc.Start()
	doc.AddInt("state", 1)
	doc.AddStr("info", "info")
	doc.AddFloat64("float", 123.123)
	doc.StartTitle("data")
	doc.AddStr("name", "wang")
	doc.End()
	doc.StartArr("list")
	doc.Start()
	doc.AddStr("type", "8888")
	doc.End()
	doc.Start()
	doc.AddStr("type", "中文")
	doc.End()
	doc.EndArr()
	doc.AddStr("time", "time")
	doc.End()
	return common.ResponseEntity{State: common.SUCCESS, Data: doc.String()}
}
