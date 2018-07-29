package table

import (
	"github.com/mingz2013/mahjong-table-go/msg"
	"github.com/mingz2013/mahjong-table-go/player"
	"log"
	"time"
)

type Table struct {
	Id string // 桌子Id，由manager自动生成

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	Play    Play
	Players [4]player.Player
}

func NewTable(id string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Table {
	t := Table{Id: id, MsgIn: msgIn, MsgOut: msgOut}
	t.Init()
	return t
}

func (t *Table) Init() {
	for i := 0; i < 4; i++ {
		t.Players[i] = player.NewPlayer(i)
	}
	t.Play = NewPlay(t)

}

func (t Table) Run() {

	for {
		select {
		case m, ok := <-t.MsgIn:
			{
				if !ok {
					continue
				}

				t.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue

		}

	}
}

func (t *Table) onMsg(m msg.Msg) {
	log.Println("Table.onMsg", m, &t, &t.Players)
	switch m.GetCmd() {
	case "table":
		t.onTableMsg(m)
	case "play":
		t.Play.OnMsg(m)
	default:
		log.Println("unknown cmd", m)
	}
}

func (t *Table) onTableMsg(m msg.Msg) {
	params := m.GetParams()
	action := params["action"].(string)
	switch action {
	case "sit":
		t.onTableSit(m)
	}

}

func (t *Table) getBestSeatId() (int, bool) {
	log.Println("getBestSeatId", "players", t.Players, &t.Players)
	for i := 0; i < 4; i++ {
		log.Println("check Id", t.Players[i].Id)
		if t.Players[i].Id == 0 {
			return i, true
		}
	}
	return -1, false
}

func (t *Table) onTableSit(m msg.Msg) {
	params := m.GetParams()
	id := params["id"].(int)
	name := params["name"].(string)

	seatId, ok := t.getBestSeatId()

	if !ok {
		log.Println("not found empty seat", id, name)
		t.SendTableSitRes(id, map[string]interface{}{"retcode": -1, "msg": "not found empty seat"})
		return
	}

	p := t.Players[seatId]
	p.Id = id
	p.SeatId = seatId
	p.Name = name

	t.Players[seatId] = p
	log.Println("players", t.Players, "p", p)

	log.Println("sit ok, ", "userId", id, "name", name, "seatId", seatId)

	t.SendTableSitRes(id, map[string]interface{}{"retcode": 0, "msg": "sit ok"})

	t.checkFull()
}

func (t *Table) SendTableSitRes(id int, results map[string]interface{}) {
	t.SendTableRes(id, "sit", results)
}

func (t *Table) SendTableRes(id int, action string, results map[string]interface{}) {
	//t.MsgOut <- msg.Msg{"id": id, "cmd": "table", "results": results}

	results["action"] = action
	t.SendRes(id, "table", results)
}

func (t *Table) SendRes(id int, cmd string, results map[string]interface{}) {
	t.MsgOut <- msg.Msg{"id": id, "cmd": cmd, "results": results}
}

func (t *Table) checkFull() {
	_, ok := t.getBestSeatId()
	if !ok {
		t.onTableStart()
	} else {
		log.Println("not all full...")
	}
}

func (t *Table) onTableStart() {
	t.Play.Run()
}
