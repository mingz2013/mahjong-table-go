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

func (p *Player) GetShadowInfo() {

}

func (p *Player) GetActionsInfo() (info []interface{}) {
	for i := 0; i < len(p.actions); i++ {
		info = append(info, p.actions[i].GetInfo())
	}
	return
}

func (p *Player) AddAction(action actions.BaseAction) {
	p.actions = append(p.actions, action)
}

func (p *Player) choiceAction() {
	// 用户选择操作，选择了一个按钮
}

func (p *Player) isChoosedAction() {
	// 用户是否选择过了操作
}

func (p *Player) getBestActionSelf() {
	// 获取自己可操作的actions中优先级最高的action
}

func (p *Player) getWorstActionSelf() {
	// 获取自己等级优先级最低的action，用于自动出牌
}

func (p *Player) clearChoosedAction() {

}

func (p *Player) isActionEmpty() {

}

func (p *Player) addActions() {

}

func (p *Player) doAction() {

}
