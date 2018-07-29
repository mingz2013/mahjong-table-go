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
		action := actions.NewKaiPaiAction(tiles)
		player_ := p.table.Players[i]
		player_.DoKaiPaiAction(action)
		//action.DoAction(player_)
		p.afterKaiPaiAction(player_, action)
	}

}

func (p *Play) afterKaiPaiAction(player player.Player, action actions.KaiPaiAction) {
	p.SendPlayKaiPaiRes(player, action)
}

func (p *Play) SendPlayKaiPaiRes(player player.Player, action actions.KaiPaiAction) {
	p.SendPlayRes(player, "kai_pai", action.GetInfo())
}

func (p *Play) nextOp(seatId int) {
	log.Println("Play.nextOp...", seatId)
	p.moPai(p.table.Players[seatId])
}

func (p *Play) moPai(player player.Player) {
	tile, ok := p.PopMoPai()
	log.Println("Play.moPai..", tile, ok)
	if !ok {
		p.onPlayFlow()
	} else {
		action := actions.NewMoPaiAction(tile)
		//action.DoAction(player)
		player.DoMoPaiAction(action)
		p.afterMoPaiAction(player, action)
	}

}

func (p *Play) afterMoPaiAction(player player.Player, action actions.MoPaiAction) {
	p.SendPlayMoPaiRes(player, action)
}

func (p *Play) SendPlayMoPaiRes(player player.Player, action actions.MoPaiAction) {
	p.SendPlayRes(player, "mo_pai", action.GetInfo())
}

func (p *Play) onPlayFlow() {

}

func (p *Play) SendPlayRes(player player.Player, action string, results map[string]interface{}) {
	results["action"] = action
	p.table.SendRes(player.Id, "play", results)
}
