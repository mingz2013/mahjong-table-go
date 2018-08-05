package table

import "github.com/mingz2013/mahjong-table-go/player"

type PlayersManager struct {
	Players [4]player.Player
}

func (p *PlayersManager) Init() {
	for i := 0; i < 4; i++ {
		p.Players[i] = player.NewPlayer(i)
	}
	//t.Play = NewPlay(t)

}

func (p *PlayersManager) isAllPlayersChoosedAction() {

}

func (p *PlayersManager) isAllPlayersActionEmpty() {

}

func (p *PlayersManager) getBestChoosedActionPlayer() {

}

func (p *PlayersManager) isPlayersBestAction() {

}

func (p *PlayersManager) isPlayersHasHuAction() {

}

func (p *PlayersManager) getAllPlayersCards() {

}
