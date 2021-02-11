package cards

import (
	"log"
	"mahjong-table-go/actions"
	"sort"
)

type Pile struct {
}

type HandPile struct {
	pattern []int
}

func (h *HandPile) getInfo() {

}

func (h *HandPile) parseFromInfo() {

}

func (h *HandPile) getShadowInfo() {

}

func (h *HandPile) addTile(tile int) {
	h.pattern = append(h.pattern, tile)
	sort.Ints(h.pattern)
	//h.pattern.PushBack(tile)
	//h.pattern
}

func (h *HandPile) rmTile(tile int) {
	for i := 0; i < len(h.pattern); i++ {
		if h.pattern[i] == tile {
			h.pattern = append(h.pattern[:i], h.pattern[i+1:]...)
			return
		}
	}
}

func (h *HandPile) rmTiles() {

}

func (h *HandPile) popTile() {

}

func (h *HandPile) isHaveTile() {

}

func (h *HandPile) countTile(tile int) {

}

func (h *HandPile) hasChi(tile int) {

}

func (h *HandPile) hasPeng(tile int) {

}

func (h *HandPile) hasDianGang(tile int) {

}

func (h *HandPile) hasAnGang() {

}

func (h *HandPile) hasHu(tile int) []interface{} {
	return []interface{}{}
}

func (h *HandPile) ChoiceTileToDrop() int {
	return h.pattern[len(h.pattern)-1]
}

type DropPile struct {
	pattern        []int // 对应前端显示的，不包括被别人拿走的
	patternHistory []int // 所有的历史记录
}

func (d *DropPile) parseFromInfo() {

}

func (d *DropPile) GetInfo() {

}

func (d *DropPile) AddTile(tile int) {
	d.pattern = append(d.pattern, tile)
}

func (d *DropPile) RmTile() {

}

func (d *DropPile) CountTile() {

}

type ChiPile struct {
	pattern    []int
	tile       int
	fromSeatId int
}

func (c *ChiPile) countTile() {

}

func (c *ChiPile) parseFromInfo() {

}

func (c *ChiPile) getInfo() {

}

func (c *ChiPile) parseFromAction() {

}

type PengPile struct {
	pattern    []int
	tile       int
	fromSeatId int
}

func (p *PengPile) parseFromInfo() {

}

func (p *PengPile) getInfo() {

}

func (p *PengPile) parseFromAction() {

}

func (p *PengPile) countTile() {

}

type GangPile struct {
	pattern    []int
	tile       int
	gangType   int
	fromSeatId int
}

func (g *GangPile) parseFromInfo() {

}

func (g *GangPile) getInfo() {

}

func (g *GangPile) getShadowInfo() {

}

func (g *GangPile) parseFromAction() {

}

type Cards struct {
	handPile HandPile
	dropPile DropPile
	cpgPiles []Pile
	nowTile  int
}

func (c *Cards) Init() {
	c.handPile = HandPile{}
	c.dropPile = DropPile{}
}

func NewCards() Cards {
	return Cards{}
}

func (c *Cards) parseFromInfo() {

}

func (c *Cards) getInfo() {

}

func (c *Cards) getShadowInfo() {

}

func (c *Cards) getNowTile() {

}

func (c *Cards) getShadowNowTile() {

}

func (c *Cards) rmDropTile() {

}

func (c *Cards) DoAction(action actions.BaseAction) {
	log.Println("Cards.DoAction...", action)

	switch action.GetName() {
	case "kai_pai":
		c.DoKaiPaiAction(action.(actions.KaiPaiAction))
	case "mo_pai":
		c.DoMoPaiAction(action.(actions.MoPaiAction))
	case "chu_pai":
		c.DoChuPaiAction(action.(actions.ChuPaiAction))
	default:
		log.Println("known action", action)

	}

}

func (c *Cards) DoKaiPaiAction(action actions.KaiPaiAction) {
	copy(c.handPile.pattern, action.Tiles)
}

func (c *Cards) DoMoPaiAction(action actions.MoPaiAction) {
	c.nowTile = action.Tile
}

func (c *Cards) DoChuPaiAction(action actions.ChuPaiAction) {
	if c.nowTile > 0 {
		c.handPile.addTile(c.nowTile)
		c.nowTile = -1
	}
	c.handPile.rmTile(action.Tile)
	c.dropPile.AddTile(action.Tile)

}

// 一些算法

func (c *Cards) ChoiceActionToDo(actions []actions.BaseAction) actions.BaseAction {
	// 选择一个操作去做
	log.Println("Cards.ChoiceActionToDo", actions)

	for i := 0; i < len(actions); i++ {
		for j := i + 1; j < len(actions); j++ {
			if actions[i].GetLevel() < actions[j].GetLevel() {
				actions[i], actions[j] = actions[j], actions[i]
			}
		}
	}

	return actions[len(actions)-1]
}

func (c *Cards) ChoiceTileToDrop() int {
	// 选一张牌打出去
	if c.nowTile > 0 {
		return c.nowTile
	}
	return c.handPile.ChoiceTileToDrop()
}

func (c *Cards) assertCount() {

}

func (c *Cards) hasChiWithChuPai() {

}

func (c *Cards) hasPengWithChuPai() {

}

func (c *Cards) hasDianGangWithChuPai() {

}

func (c *Cards) hasAnGangWithMoPai() {

}

func (c *Cards) hasBuGangWithMoPai() {

}

func (c *Cards) hasHuWithMoPai() []interface{} {
	return c.handPile.hasHu(c.nowTile)
}

func (c *Cards) hasHuWithChuPai(tile int) []interface{} {
	return c.handPile.hasHu(tile)
}

func (c *Cards) hasHuWithGangPai() {

}
