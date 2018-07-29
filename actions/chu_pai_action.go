package actions

type ChuPaiAction struct {
	tile int
	name string
}

func NewChuPaiAction(tile int) ChuPaiAction {
	return ChuPaiAction{tile, "chu_pai"}
}

func (a *ChuPaiAction) ParseFromInfo(action map[string]interface{}) ChuPaiAction {
	a.tile = action["tile"].(int)
	a.name = "chu_pai"
	return *a
}

func (a *ChuPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tile": a.tile,
		"name": a.name,
	}
}

func (a *ChuPaiAction) GetName() string {
	return a.name
}
