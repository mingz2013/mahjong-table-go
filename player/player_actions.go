package player

import (
	"log"
	"mahjong-table-go/actions"
)

type PlayerActions struct {
	actions []actions.BaseAction

	ChoosedAction actions.BaseAction
}

func (p *PlayerActions) Init() {
	p.actions = []actions.BaseAction{}
	p.ChoosedAction = nil
}

func (p *PlayerActions) GetActionsInfo() (info []interface{}) {
	for i := 0; i < len(p.actions); i++ {
		info = append(info, p.actions[i].GetInfo())
	}
	return
}

func (p *PlayerActions) AddActions(actions []actions.BaseAction) {
	for i := 0; i < len(actions); i++ {
		p.AddAction(actions[i])
	}
}

func (p *PlayerActions) AddAction(action actions.BaseAction) {
	p.actions = append(p.actions, action)
}

func (p *PlayerActions) RemoveAction(action actions.BaseAction) {
	for i := 0; i < len(p.actions); i++ {
		if action.IsValid(p.actions[i]) {
			p.actions = append(p.actions[:i], p.actions[i+1:]...)
			break
		}
	}
}

func (p *PlayerActions) ClearActions() {
	p.actions = []actions.BaseAction{}
}

func (p *PlayerActions) IsCanDoAction(action actions.BaseAction) (bool, actions.BaseAction) {
	log.Println("PlayerActions.IsCanDoAction", action)
	for i := 0; i < len(p.actions); i++ {
		if action.IsValid(p.actions[i]) {

			action.UpdateLocalAction(p.actions[i])

			return true, action
		}
	}

	return false, action
}

func (p *PlayerActions) ChoiceAction(action actions.BaseAction) bool {
	log.Println("PlayerActions.ChoiceAction", action)

	// 用户选择操作，选择了一个按钮
	isCanDo, action := p.IsCanDoAction(action)
	if isCanDo {
		p.ChoosedAction = action
		p.RemoveAction(action)
		p.ClearActions()
		return true
	}

	return false
}

func (p *PlayerActions) IsChoosedAction() bool {
	// 用户是否选择过了操作
	return len(p.actions) == 0
}

func (p *PlayerActions) GetBestActionSelf() actions.BaseAction {
	// 获取自己可操作的actions中优先级最高的action
	return nil
}

func (p *PlayerActions) GetWorstActionSelf() {
	// 获取自己等级优先级最低的action，用于自动出牌
}

func (p *PlayerActions) ClearChoosedAction() {

}

func (p *PlayerActions) IsActionEmpty() bool {
	//
	return len(p.actions) == 0
}
