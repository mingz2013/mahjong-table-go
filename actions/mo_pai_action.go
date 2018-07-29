package actions

type MoPaiAction struct {
	Tile int
}

func NewMoPaiAction(tile int) MoPaiAction {
	return MoPaiAction{tile}
}

func (a *MoPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tile": a.Tile,
	}
}
