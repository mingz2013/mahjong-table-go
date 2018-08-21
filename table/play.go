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
	p.xiPai()
}

func NewPlay(t *Table) Play {
	p := Play{table: t}
	p.Init()
	return p
}

func (p *Play) GetInfo() {

}

func (p *Play) ChangeCurrentSeatId() {

}

func (p *Play) GetNextSeatId() {

}

func (p *Play) OnMsg(m msg.Msg) {
	log.Println("Play.OnMsg->", m)
	action := m.GetParams()["action"].(string)
	userId := m.GetKey("id").(int)

	_player := p.table.GetPlayersByUserId(userId)
	var _action actions.BaseAction
	switch action {
	case "chu_pai":
		chuPaiAction := actions.ChuPaiAction{}
		chuPaiAction.ParseFromInfo(m.GetParams())
		_action = chuPaiAction
	default:
		log.Println("Play.OnMsg..err.....", m)
	}

	_player.ChoiceAction(_action)
	p.DoActionCheckToDo(_player, _action)
}

func (p *Play) DoActionCheckToDo(player *player.Player, action actions.BaseAction) {
	// 所有人都选择了, 选择其中最高优先级的
	// 当前选择的是未选择中玩家最高优先级的
	// 之前选择过的，已经有最高优先级的了
	if p.table.isAllPlayersChoosedAction() || p.table.isPlayersBestAction(player, action) || p.table.isBestActionComming() {
		bestPlayer := p.table.GetBestChoosedActionPlayer()
		bestAction := bestPlayer.ChoosedAction

		bestPlayer.DoAction(bestAction)

	}

}

func (p *Play) Run() {
	log.Println("play run...")
	p.start()
}

func (p *Play) start() {
	log.Println("Play.start...players", p.table.Players)
	p.kaiPai()
	p.nextOp(0)
}

func (p *Play) kaiPai() {
	log.Println("Play.kaiPai...")
	for i := 0; i < 4; i++ {
		tiles := p.PopKaiPai()
		action := actions.NewKaiPaiAction(tiles)
		player_ := p.table.Players[i]

		log.Println("Play.kaiPai..", "p.table.Players", p.table.Players)

		player_.DoKaiPaiAction(action)
		//action.DoAction(player_)

		log.Println("Play.kaiPai..", player_, action)

		p.afterActionKaiPai(player_, action)
	}

}

func (p *Play) SendPlayKaiPaiRes(player *player.Player, action actions.KaiPaiAction) {

	results := action.GetInfo()
	results["actions"] = player.GetActionsInfo()

	p.SendPlayRes(player, "kai_pai", results)
}

func (p *Play) nextOp(seatId int) {
	log.Println("Play.nextOp...", seatId)
	p.moPai(p.table.Players[seatId])
}

func (p *Play) moPai(player *player.Player) {
	tile, ok := p.PopMoPai()
	log.Println("Play.moPai..", tile, ok)
	if !ok {
		p.onPlayFlow()
	} else {
		action := actions.NewMoPaiAction(tile)
		//action.DoAction(player)
		player.DoMoPaiAction(action)
		p.afterActionMoPai(player, action)
	}

}

func (p *Play) afterActionMoPai(player *player.Player, action actions.MoPaiAction) {
	chuPaiAction := actions.NewChuPaiAction(-1)
	player.AddAction(&chuPaiAction)

	p.SendPlayMoPaiRes(player, action)
}

func (p *Play) afterActionKaiPai(player *player.Player, action actions.KaiPaiAction) {
	log.Println("Play.afterActionKaiPai", "player", player, "action", action)
	p.SendPlayKaiPaiRes(player, action)
}

func (p *Play) SendPlayMoPaiRes(player *player.Player, action actions.MoPaiAction) {
	results := action.GetInfo()
	results["actions"] = player.GetActionsInfo()
	p.SendPlayRes(player, "mo_pai", results)
}

func (p *Play) onPlayFlow() {

}

func (p *Play) onPlayWin() {

}

func (p *Play) onPlayOver() {

}

func (p *Play) SendPlayRes(player *player.Player, action string, results map[string]interface{}) {
	results["action"] = action

	p.table.SendRes(player.Id, "play", results)
}
