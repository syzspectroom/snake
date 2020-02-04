package game

import "errors"

func nToXY(n uint) (x uint, y uint, err error) {
	if n > WIDTH*HEIGHT-1 {
		return 0, 0, errors.New("board size is smaller")
	}
	return n - ((n / WIDTH) * WIDTH), n / WIDTH, nil
}

func xyToN(x uint, y uint) (n uint, err error) {
	if x > WIDTH-1 || y > HEIGHT-1 {
		return 0, errors.New("board size is smaller")
	}
	return y*WIDTH + x, nil
}
