package game

import (
	"fmt"
	"math/rand"
)

const (
	// board size
	WIDTH  = 32
	HEIGHT = 15
	// board cell values
	FLOOR = 0
	WALL  = 1
)

type Game struct {
	Board   [WIDTH * HEIGHT]uint8
	Bonuses map[*Bonus]bool
	Snakes  map[*Snake]bool
}

func NewGame() *Game {
	g := &Game{
		Bonuses: make(map[*Bonus]bool),
		Snakes:  make(map[*Snake]bool),
	}
	g.initBoard()
	return g
}

func (g *Game) initBoard() {
	for n := uint(0); n < WIDTH*HEIGHT; n++ {
		g.Board[n] = g.initCell(n)
	}
	g.addBonuses()
}

func (g *Game) initCell(n uint) uint8 {
	switch {
	case n >= 0 && n < WIDTH:
		return WALL
	case n < WIDTH*HEIGHT && n >= WIDTH*HEIGHT-WIDTH:
		return WALL
	case n%WIDTH == 0:
		return WALL
	case n%WIDTH == WIDTH-1:
		return WALL
	default:
		return FLOOR
	}
}

func (g *Game) addBonuses() {
	bonusesCount := rand.Intn(10)
	for i := 0; i < bonusesCount; i++ {
		g.addRandomCherry()
	}
}

func (g *Game) addRandomCherry() {
	// TODO: add weights to generate more balanced values
	x := uint(rand.Intn(WIDTH))
	y := uint(rand.Intn(HEIGHT))

	n, err := xyToN(x, y)
	if err != nil {
		g.addRandomCherry()
	}

	if g.isEmpty(n) {
		bonus := newBonus(n, CHERRY, 0)
		g.Bonuses[bonus] = true
	} else {
		g.addRandomCherry()
	}
}

func (g *Game) AddSnake(snake *Snake) {
	n, _ := xyToN(1, HEIGHT/2)
	snake.Place(n + HEIGHT*uint(len(g.Snakes)))
	g.Snakes[snake] = true
}

func (g *Game) isEmpty(n uint) bool {
	return g.getCell(n) == FLOOR
}

func (g *Game) getCell(n uint) uint8 {
	return g.Board[n]
}

func (g Game) ToArray() [WIDTH * HEIGHT]uint8 {

	for bonus := range g.Bonuses {
		g.Board[bonus.pos] = bonus.val
	}

	for snake := range g.Snakes {
		s := snake.Print()
		for n, v := range *s {
			g.Board[n] = uint8(v)
		}
	}
	return g.Board
}

func (g *Game) Tick() {
	for snake := range g.Snakes {
		snake.Tick()
	}

	g.calculateCollisions()
}

func (g Game) calculateCollisions() {
	// TODO: Create array of all snakes positions to calculate positions
	// var heads []uint
	// var bodies []uint
	// for snake := range g.Snakes {
	// 	if !g.isEmpty(snake.Head) {

	// 	}

	// 	for collisionSnake := range g.Snakes {

	// 	}
	// 	delete(g.Snakes, snake)
	// }
	// 	snake.head = 1
	// 	collicollisionSnake.BodyTiles
	// }

	// }
	// log.Printf("bodies: %+v \n", bodies)

}

func (g *Game) Debug() {
	z := g.ToArray()
	for y := uint(0); y < HEIGHT; y++ {
		for x := uint(0); x < WIDTH; x++ {
			n, _ := xyToN(x, y)
			fmt.Printf("%+v", z[n])
		}
		fmt.Printf("\n")
	}
}
