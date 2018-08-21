package actions

type ChuPaiAction struct {
	Tile  int
	name  string
	level int
}

func NewChuPaiAction(tile int) ChuPaiAction {
	return ChuPaiAction{tile, "chu_pai", ACTION_LEVEL_CHU_PAI}
}

func (a *ChuPaiAction) ParseFromInfo(action map[string]interface{}) ChuPaiAction {
	a.Tile = action["tile"].(int)
	a.name = "chu_pai"
	return *a
}

func (a ChuPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tile":  a.Tile,
		"name":  a.name,
		"level": a.level,
	}
}

func (a ChuPaiAction) GetName() string {
	return a.name
}

func (a ChuPaiAction) GetLevel() int {
	return a.level
}

func (a ChuPaiAction) IsValid(action *BaseAction) bool {
	return false
}

func (a ChuPaiAction) UpdateLocalAction(action *BaseAction) {

}
