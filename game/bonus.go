package game

const (
	CHERRY = 5
	RABBIT = 10
)

type Bonus struct {
	pos   uint
	val   uint8
	tickN uint //tick apperiance n
}

func newBonus(pos uint, val uint8, tickN uint) *Bonus {
	return &Bonus{
		pos:   pos,
		val:   val,
		tickN: tickN}
}

func (b *Bonus) tick() {

}
