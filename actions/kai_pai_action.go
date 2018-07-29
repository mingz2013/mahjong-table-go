package actions

type KaiPaiAction struct {
	Tiles []int
}

func NewKaiPaiAction(tiles []int) KaiPaiAction {
	return KaiPaiAction{tiles}
}

func (a *KaiPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tiles": a.Tiles,
	}
}
