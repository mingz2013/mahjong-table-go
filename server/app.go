package server

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/net_base"
	"github.com/mingz2013/mahjong-table-go/manager"
	"log"
)

// 桌子进程，很有可能，只是个客户端，连接到中心服务器，不用server监听

type TableApp struct {
	net_base.Handler

	manager manager.Manager
}

func (a *TableApp) Init() {
	a.manager = manager.NewTableManager("111")
}

func NewTableApp() *TableApp {
	a := &TableApp{}
	a.Init()
	return a
}

func (a *TableApp) Serve(c net_base.Conn, buf []byte) {
	// 解析成json，ServeJson
	var js map[string]interface{}
	err := json.Unmarshal(buf, js)
	if err == nil {
		a.ServeJson(c, js)
	} else {
		log.Println(err, buf)
	}
}

func (a *TableApp) ServeJson(c net_base.Conn, js map[string]interface{}) {
	// 前端发第一个协议，bind_user, 绑定用户连接，前端数据中应该有userId和token
	//cmd := js["cmd"].(string)
	//userId := js["userId"].(int)
	//token := js["token"].(string)
	//// 验证token和userId

	a.manager.MsgIn <- js

}
