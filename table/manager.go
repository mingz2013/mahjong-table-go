package table

import (
	"github.com/mingz2013/mahjong-table-go/msg"
	"log"
)

type Manager struct {
	Id       string
	tableMap map[string]Table

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func NewTableManager(id string) Manager {
	return Manager{Id: id, tableMap: make(map[string]Table)}
}

func (mgr *Manager) Run() {

	for {
		m, ok := <-mgr.MsgIn
		if !ok {
			continue
		}
		mgr.onMsg(m)
	}

}

func (mgr *Manager) onMsg(m msg.Msg) {

	userId := m["userId"].(int)

	switch m.GetCmd() {
	case "create":
		mgr.onCmdCreate(m, userId)
	case "join":
		mgr.onCmdJoin(m, userId)
	case "table":
		// 通过session找到对应的table，然后调用table的onMsg
		break
	default:
		log.Println("unknown cmd", m)
	}

}

func (mgr *Manager) onCmdCreate(m msg.Msg, userId int) {

}

func (mgr *Manager) onCmdJoin(m msg.Msg, userId int) {

}
