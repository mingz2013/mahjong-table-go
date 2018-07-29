package player

import (
	"github.com/mingz2013/mahjong-table-go/actions"
	"github.com/mingz2013/mahjong-table-go/cards"
)

type Player struct {
	Id   int
	Name string

	SeatId int

	cards.Cards

	actions []actions.BaseAction
}

func (p *Player) Init() {
	p.Cards = cards.NewCards()
}

func NewPlayer(seatId int) Player {

	p := Player{SeatId: seatId}
	p.Init()
	return p
}

func (p *Player) GetInfo() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *Player) GetActionsInfo() (info []interface{}) {
	for i := 0; i < len(p.actions); i++ {
		info = append(info, p.actions[i].GetInfo())
	}
	return
}
