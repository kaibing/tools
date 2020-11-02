package main

import (
	"ci_client/action"
	"ci_client/common"
	"ci_client/log"
	"container/list"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	port := common.GetPort()

	// 初始化服务器
	mux := http.NewServeMux()
	mux.Handle("/", cmd)
	log.Info.Println("listen port", port)
	_ = http.ListenAndServe(port, mux)
}

type command struct {
	t          reflect.Type
	methodInfo map[string]*methodInfo
	mapper     map[string]interface{}
	action     *list.List
}

type methodInfo struct {
	command interface{}
	method  *reflect.Method
}

// 单例对象
var cmd *command

// 初始化
func init() {
	if cmd == nil {
		cmd = &command{}
	}
	v := reflect.TypeOf(cmd)
	cmd.t = v
}

// 注册action
func init() {
	cmd.action = list.New()
	cmd.action.PushBack(&action.ClientAction{})
}

func init() {
	cmd.mapper = make(map[string]interface{})
	for clientAction := cmd.action.Front(); clientAction != nil; clientAction = clientAction.Next() {
		t := reflect.TypeOf(clientAction.Value).Elem()
		var mpg string
		field, b := t.FieldByName("path")
		if !b {
			mpg = t.Name()
			continue
		} else {
			mpg = string(field.Tag)
		}
		cmd.mapper[mpg] = clientAction.Value
	}
}

func init() {
	cmd.methodInfo = make(map[string]*methodInfo)
	for k, a := range cmd.mapper {
		v := reflect.TypeOf(a)
		for i := 0; i < v.NumMethod(); i++ {
			info := methodInfo{}
			method := v.Method(i)
			name := method.Name
			cmd.methodInfo[k+"/"+name] = &info
			info.command = a
			info.method = &method
		}
	}
}

func (cmd *command) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info.Println(*r.URL)
	// handle error
	defer func() {
		if err := recover(); err != nil {
			panic(err)
			log.Error.Println(err)
			cmd.doResponse(w, common.Failed)
		}
	}()
	// find method
	method := cmd.findMethod(r)
	if method == nil {
		log.Error.Println("not find", *r.URL)
		cmd.doResponse(w, common.Failed)
		return
	}
	// parse param
	param := cmd.parseParam(r)

	// response
	values := []reflect.Value{reflect.ValueOf(method.command), reflect.ValueOf(common.ParamEntity{param})}
	call := method.method.Func.Call(values)
	response := cmd.handleResponse(w, call)
	cmd.doResponse(w, response)
}

// 查找对应的方法
func (cmd *command) findMethod(request *http.Request) *methodInfo {
	path := request.URL.Path
	methodName := path[1:]
	return cmd.methodInfo[methodName]
}

// 解析参数
func (cmd *command) parseParam(request *http.Request) map[string]string {
	value := request.URL.Query()
	bodyByte, _ := ioutil.ReadAll(request.Body)

	res := make(map[string]string)
	_ = json.Unmarshal(bodyByte, &res)
	for k, v := range value {
		res[k] = v[0]
	}
	return res
}

func (cmd *command) handleResponse(writer http.ResponseWriter, call []reflect.Value) common.ResponseEntity {
	value := call[0]
	entity := value.Interface().(common.ResponseEntity)
	return entity
}

func (cmd *command) doResponse(w http.ResponseWriter, response common.ResponseEntity) {
	_, err := w.Write(response.GetBytes())
	if err != nil {
		log.Error.Println(err)
	}
}
