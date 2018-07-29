package actions

type KaiPaiAction struct {
	Tiles []int
	name  string
}

func NewKaiPaiAction(tiles []int) KaiPaiAction {
	return KaiPaiAction{tiles, "kai_pai"}
}

func (a *KaiPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tiles": a.Tiles,
		"name":  a.name,
	}
}

func (a *KaiPaiAction) GetName() string {
	return a.name
}
