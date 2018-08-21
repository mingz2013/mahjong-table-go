package player

import "github.com/mingz2013/mahjong-table-go/actions"

type PlayerActions struct {
	actions []actions.BaseAction

	ChoosedAction actions.BaseAction
}

func (p *PlayerActions) GetActionsInfo() (info []interface{}) {
	for i := 0; i < len(p.actions); i++ {
		info = append(info, p.actions[i].GetInfo())
	}
	return
}

func (p *PlayerActions) AddAction(action actions.BaseAction) {
	p.actions = append(p.actions, action)
}

func (p *PlayerActions) RemoveAction(action actions.BaseAction) {
	for i := 0; i < len(p.actions); i++ {

	}
}

func (p *PlayerActions) ClearActions() {
	p.actions = []actions.BaseAction{}
}

func (p *PlayerActions) IsCanDoAction(action actions.BaseAction) (bool, actions.BaseAction) {
	return false, action
}

func (p *PlayerActions) ChoiceAction(action actions.BaseAction) bool {
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

func (p *PlayerActions) IsActionEmpty() {

}

func (p *PlayerActions) AddActions() {

}
