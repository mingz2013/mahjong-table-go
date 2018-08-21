package actions

type MoPaiAction struct {
	Tile  int
	name  string
	level int
}

func NewMoPaiAction(tile int) MoPaiAction {
	return MoPaiAction{tile, "mo_pai", ACTION_LEVEL_MO_PAI}
}

func (a MoPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tile":  a.Tile,
		"name":  a.name,
		"level": a.level,
	}
}

func (a MoPaiAction) GetName() string {
	return a.name
}

func (a MoPaiAction) GetLevel() int {
	return a.level
}

func (a MoPaiAction) IsValid(action *BaseAction) bool {
	return false
}

func (a MoPaiAction) UpdateLocalAction(action *BaseAction) {

}
