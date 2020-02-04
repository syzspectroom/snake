package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"snake/game"
	"snake/transport"
)

var (
	server   = transport.NewServer()
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func ws(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	server.AddClient(conn)
	return nil
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", ws)
	e.Static("/", "./public")

	go broadcast()

	a := game.NewGame()

	s := game.NewSnake(1)
	b := game.NewSnake(2)
	a.AddSnake(s)
	a.AddSnake(b)

	a.Debug()
	s.NextMove(game.MOVE_N)
	a.Tick()
	a.Debug()
	a.Tick()
	a.Debug()

	s.NextMove(game.MOVE_W)
	a.Tick()
	a.Debug()
	a.Tick()
	a.Debug()

	s.NextMove(game.MOVE_E)
	a.Tick()
	a.Debug()

	e.Logger.Fatal(e.Start(":1323"))

}

func broadcast() {
	for {
		time.Sleep(5 * time.Second)
		server.Deliver([]byte("test"))
	}
}
