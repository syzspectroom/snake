package game

const (
	SNAKE_START_SIZE = 3

	//body codes
	HEAD = 8
	BODY = 7
	//move direction constants
	MOVE_W = 1
	MOVE_E = 2
	MOVE_S = 3
	MOVE_N = 4
)

type Snake struct {
	color         uint
	size          uint
	moveDirection uint8
	Head          uint   //head position
	BodyTiles     []uint //body positions array
}

func NewSnake(color uint) *Snake {
	return &Snake{
		color:         color,
		size:          SNAKE_START_SIZE,
		moveDirection: MOVE_E,
	}
}

//Place - place the snake on a board. n argument is a tail position
func (s *Snake) Place(n uint) {
	s.Head = n + (s.size - 1)

	for x := n; x < s.Head; x++ {
		s.BodyTiles = append(s.BodyTiles, x)
	}
}

func (s *Snake) Print() *map[uint]int {
	a := make(map[uint]int)
	for _, x := range s.BodyTiles {
		a[x] = BODY
	}
	a[s.Head] = HEAD

	return &a
}

func (s *Snake) CheckCollide(n uint) bool {
	if s.Head == n {
		return true
	}
	for i := uint(0); i < s.size; i++ {
		if s.BodyTiles[i] == n {
			return true
		}
	}
	return false
}

func (s *Snake) NextMove(direction uint8) {
	// Can`t move backwards
	if (direction == MOVE_E && s.moveDirection == MOVE_W) ||
		(direction == MOVE_W && s.moveDirection == MOVE_E) ||
		(direction == MOVE_S && s.moveDirection == MOVE_N) ||
		(direction == MOVE_N && s.moveDirection == MOVE_S) {
		s.moveDirection = s.moveDirection
	} else {
		s.moveDirection = direction
	}
}

func (s *Snake) performMove() {
	s.BodyTiles = append(s.BodyTiles[1:], s.Head)
	switch s.moveDirection {
	case MOVE_W:
		s.Head--
	case MOVE_E:
		s.Head++
	case MOVE_S:
		s.Head = s.Head + WIDTH
	case MOVE_N:
		s.Head = s.Head - WIDTH
	}
}

func (s *Snake) Tick() {
	s.performMove()
}
