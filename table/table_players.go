package table

import (
	"github.com/mingz2013/mahjong-table-go/actions"
	"github.com/mingz2013/mahjong-table-go/player"
)

type TablePlayers struct {
	Players [4]*player.Player
}

func (p *TablePlayers) Init() {
	for i := 0; i < 4; i++ {
		p.Players[i] = player.NewPlayer(i)
	}
	//t.Play = NewPlay(t)

}

func (p *TablePlayers) isAllPlayersChoosedAction() bool {
	// 是否所有用户都选择了action，
	for i := 0; i < len(p.Players); i++ {
		if !p.Players[i].IsChoosedAction() {
			return false
		}
	}
	return true
}

func (p *TablePlayers) isAllPlayersActionEmpty() {

}

func (p *TablePlayers) GetBestChoosedActionPlayer() (bestPlayer *player.Player) {
	for i := 0; i < len(p.Players); i++ {
		player_ := p.Players[i]
		if player_.ChoosedAction != nil {
			if bestPlayer == nil {
				bestPlayer = player_
				continue
			}

			if player_.ChoosedAction.GetLevel() > bestPlayer.ChoosedAction.GetLevel() {
				bestPlayer = player_

			} else if player_.ChoosedAction.GetLevel() == bestPlayer.ChoosedAction.GetLevel() {
				bestPlayer = player_
			} else {

			}

		}
	}

	return
}

func (p *TablePlayers) isPlayersBestAction(player *player.Player, action actions.BaseAction) bool {
	// 是否是所有玩家的action中最高优先级的

	for i := 0; i < len(p.Players); i++ {
		player_ := p.Players[i]

		if player_ == player {
			continue
		}

		if player_.IsChoosedAction() {
			if player_.ChoosedAction != nil {
				if player_.ChoosedAction.GetLevel() > action.GetLevel() {
					return false
				} else if player_.ChoosedAction.GetLevel() == action.GetLevel() {
					//
					return false
				}
			}
		} else {
			p_action := player_.GetBestActionSelf()
			if p_action.GetLevel() > action.GetLevel() {
				return false
			} else if p_action.GetLevel() == action.GetLevel() {
				return false
			}
		}

	}

	return true
}

func (p *TablePlayers) isPlayersHasHuAction() {

}

func (p *TablePlayers) getAllPlayersCards() {

}

func (p *TablePlayers) isBestActionComming() bool {
	// 最高优先级的action出现了
	for i := 0; i < len(p.Players); i++ {
		_player := p.Players[i]
		if _player.IsChoosedAction() {
			if _player.ChoosedAction != nil {
				if p.isPlayersBestAction(_player, _player.ChoosedAction) {
					return true
				}
			}
		}
	}

	return false
}

func (p *TablePlayers) GetPlayersByUserId(userId int) *player.Player {
	for i := 0; i < len(p.Players); i++ {
		if p.Players[i].Id == userId {
			return p.Players[i]
		}
	}
	return nil
}
