package robot

import (
	"github.com/mingz2013/mahjong-table-go/msg"
	"github.com/mingz2013/mahjong-table-go/player"
	"log"
	"os"
	"time"
)

type Robot struct {
	Id     int
	Name   string
	Cards  player.Cards
	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func (r *Robot) Init() {
	r.Cards = player.NewCards()
}

func NewRobot(id int, name string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Robot {
	r := Robot{Id: id, Name: name, MsgIn: msgIn, MsgOut: msgOut}
	r.Init()
	return r
}

func (r *Robot) doSit() {

	r.SendTableSitReq(map[string]interface{}{"id": r.Id, "name": r.Name})

}

func (r *Robot) SendTableSitReq(params map[string]interface{}) {
	r.SendTableReq("sit", params)
}

func (r *Robot) SendTableReq(action string, params map[string]interface{}) {
	params["action"] = action
	r.SendReq("table", params)
}

func (r *Robot) SendReq(cmd string, params map[string]interface{}) {
	r.MsgOut <- msg.Msg{"cmd": cmd, "params": params}
}

func (r Robot) Run() {

	r.doSit()

	for {

		select {
		case m, ok := <-r.MsgIn:
			{
				if !ok {
					continue
				}

				r.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue
		}

	}
}

func (r *Robot) onMsg(m msg.Msg) {
	switch m.GetCmd() {
	case "table":
		r.onTableMsg(m)
	case "play":
		r.onPlayMsg(m)
	default:
		log.Println("unknown msg", m)

	}
}

func (r *Robot) onTableMsg(m msg.Msg) {

	results := m.GetResults()
	action := results["action"].(string)

	switch action {
	case "sit":
		r.onTableSitMsg(m)
	default:
		log.Println(m)
	}

}

func (r *Robot) onTableSitMsg(m msg.Msg) {
	results := m.GetResults()
	retCode := results["retcode"].(int)
	msgRet := results["msg"].(string)
	if retCode != 0 {
		log.Println(r, msgRet)
		os.Exit(retCode)
	}
	log.Println(r, msgRet)
}

func (r *Robot) onPlayMsg(m msg.Msg) {
	results := m.GetResults()
	action := results["action"].(string)

	switch action {
	case "kai_pai":
		break
	default:
		log.Println(m)
	}

}
