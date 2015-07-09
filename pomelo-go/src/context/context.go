// author:李为胜 2015-7-7

/*
package context用于记录一些全局的信息，包括：注册的modules，master服务器的配置信息，servers的配置信息
*/
package context

import (
	"fmt"
	"module"
	"os"
)

type Context struct {
	Ch         chan int8
	Env        string
	Modules    []module.Module
	MasterInfo map[string]interface{}
	ServerInfo map[string][]map[string]interface{}
}

//创建新的上下文
func NewContext() *Context {
	ch := make(chan int8)
	mods := make([]module.Module, 0, 10)
	masterInfo := make(map[string]interface{})
	serverInfo := make(map[string][]map[string]interface{})
	return &Context{ch, "", mods, masterInfo, serverInfo}
}

//向上下文中注册一个module
func (ctx *Context) RegisteModule(mod module.Module) {
	ctx.Modules = append(ctx.Modules, mod)
	fmt.Println(len(ctx.Modules))
}

// func (ctx *context) SetMasterInfo(mi map[string]interface{}) {
// 	ctx.MasterInfo = mi
// }

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error Exit: %s\n", err.Error())
		os.Exit(1)
	}
}