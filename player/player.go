package player

type Player struct {
	Id   int
	Name string

	SeatId int

	Cards
}

func (p Player) Init() {
	p.Cards = NewCards()
}

func NewPlayer(seatId int) Player {

	p := Player{SeatId: seatId}
	p.Init()
	return p
}
