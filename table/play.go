package table

import (
	"github.com/mingz2013/mahjong-table-go/actions"
	"github.com/mingz2013/mahjong-table-go/msg"
	"github.com/mingz2013/mahjong-table-go/player"
	"log"
)

type Play struct {
	table *Table
	Bottom
}

func (p *Play) Init() {
	p.initTilePool()
}

func NewPlay(t *Table) Play {
	p := Play{table: t}
	p.Init()
	return p
}

func (p *Play) OnMsg(m msg.Msg) {

}

func (p *Play) Run() {
	log.Println("play run...")
	p.start()
}

func (p *Play) start() {
	p.kaiPai()
	p.nextOp(0)
}

func (p *Play) kaiPai() {

	for i := 0; i < 4; i++ {
		tiles := p.PopKaiPai()
		kaiPaiAction := actions.NewKaiPaiAction(tiles)
		player_ := p.table.Players[i]
		player_.DoKaiPaiAction(kaiPaiAction)
		p.SendPlayKaiPaiRes(player_, kaiPaiAction)
	}

}

func (p *Play) SendPlayKaiPaiRes(player player.Player, action actions.KaiPaiAction) {
	p.SendPlayRes(player, "kai_pai", action.GetInfo())
}

func (p *Play) SendPlayRes(player player.Player, action string, results map[string]interface{}) {
	results["action"] = action
	p.table.SendRes(player.Id, "play", results)
}

func (p *Play) nextOp(seatId int) {

}
