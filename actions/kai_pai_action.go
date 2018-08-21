package actions

type KaiPaiAction struct {
	Tiles []int
	name  string
	level int
}

func NewKaiPaiAction(tiles []int) KaiPaiAction {
	return KaiPaiAction{tiles, "kai_pai", ACTION_LEVEL_KAI_PAI}
}

func (a *KaiPaiAction) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"tiles": a.Tiles,
		"name":  a.name,
		"level": a.level,
	}
}

func (a *KaiPaiAction) GetName() string {
	return a.name
}

func (a *KaiPaiAction) GetLevel() int {
	return a.level
}
