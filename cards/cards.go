package cards

import "github.com/mingz2013/mahjong-table-go/actions"

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

func (h *HandPile) addTile() {

}

func (h *HandPile) rmTile() {

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

type DropPile struct {
	pattern        []int // 对应前端显示的，不包括被别人拿走的
	patternHistory []int // 所有的历史记录
}

func (d *DropPile) parseFromInfo() {

}

func (d *DropPile) GetInfo() {

}

func (d *DropPile) AddTile() {

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

}

func (c *Cards) DoKaiPaiAction(action actions.KaiPaiAction) {
	copy(c.handPile.pattern, action.Tiles)
}

func (c *Cards) DoMoPaiAction(action actions.MoPaiAction) {
	c.nowTile = action.Tile
}

// 一些算法

func (c *Cards) ChoiceActionToDo(actions []actions.BaseAction) actions.BaseAction {
	return nil
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
