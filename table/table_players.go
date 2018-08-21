package table

import "github.com/mingz2013/mahjong-table-go/player"

type TablePlayers struct {
	Players [4]*player.Player
}

func (p *TablePlayers) Init() {
	for i := 0; i < 4; i++ {
		p.Players[i] = player.NewPlayer(i)
	}
	//t.Play = NewPlay(t)

}

func (p *TablePlayers) isAllPlayersChoosedAction() {

}

func (p *TablePlayers) isAllPlayersActionEmpty() {

}

func (p *TablePlayers) getBestChoosedActionPlayer() {

}

func (p *TablePlayers) isPlayersBestAction() {

}

func (p *TablePlayers) isPlayersHasHuAction() {

}

func (p *TablePlayers) getAllPlayersCards() {

}
