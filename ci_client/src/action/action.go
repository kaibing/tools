package action

import (
	"ci_client/common"
	"ci_client/service"
)

type ClientAction struct {
	path string `clientAction`
}

func (cmd *ClientAction) Info(param common.ParamEntity) (res common.ResponseEntity) {
	// 参数处理
	return service.GetInstance().GetInfo()
}

func (cmd *ClientAction) Start(param common.ParamEntity) (res common.ResponseEntity) {
	// 参数处理
	evnConf := param.GetStr("envConf")
	buildConf := param.GetStr("buildConf")
	id := param.GetInt("id")
	return service.GetInstance().Start(evnConf, buildConf, id)
}

func (cmd *ClientAction) Stop(param common.ParamEntity) (res common.ResponseEntity) {
	// 参数处理
	id := param.GetInt("id")
	return service.GetInstance().Stop(id)
}

func (cmd *ClientAction) Download(param common.ParamEntity) (res common.ResponseEntity) {
	// 参数处理
	return service.GetInstance().Download()
}
