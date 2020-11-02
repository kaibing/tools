package common

import (
	"strconv"
)

type ParamEntity struct {
	Params map[string]string
}

type ResponseEntity struct {
	State int
	Data  string
}

// 状态
var (
	SUCCESS = 1
	FAILED  = 2
)

// 失败相应
var (
	Failed = ResponseEntity{State: FAILED, Data: "E10010"}
	OK     = ResponseEntity{State: SUCCESS, Data: ""}
)

func (response ResponseEntity) GetBytes() []byte {
	doc := New()
	doc.Start()
	doc.AddInt("state", response.State)
	doc.AddEntity("data", response.Data)
	doc.End()
	return doc.Bytes()
}

func (param ParamEntity) GetInt(k string) int {
	r := "0"
	v := false
	if r, v = param.Params[k]; !v {
		return 0
	}
	i, err := strconv.ParseInt(r, 10, 32)
	if err != nil {
		return 0
	}
	return int(i)
}

func (param ParamEntity) GetStr(k string) string {
	r := ""
	v := false
	if r, v = param.Params[k]; !v {
		return r
	}
	return r
}

func Success(data string) (response *ResponseEntity) {
	return &ResponseEntity{State: SUCCESS, Data: data}
}
