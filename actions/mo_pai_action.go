package actions

type MoPaiAction struct {
	Tile int
	name string
}

func NewMoPaiAction(tile int) MoPaiAction {
	return MoPaiAction{tile, "mo_pai"}
}

func (a *MoPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tile": a.Tile,
		"name": a.name,
	}
}

func (a *MoPaiAction) GetName() string {
	return a.name
}
