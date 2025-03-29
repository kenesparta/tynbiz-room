package server

import (
	"flag"
	"os"
)

var (
	addr = flag.String("addr", ":", os.Getenv("PORT"), "http service address")
	cert = flag.String("cert", "", "cert file")
	key  = flag.String("key", "", "key file")
)

func Run() error {
	flag.Parse()
	if *addr == ":" {
		*addr = ":8080"
	}

	app.Get("/", handlers.Welcome)
	app.Post("/room", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/ws")
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/ws", ws.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/ws", ws.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/ws")
	app.Get("/stream/:ssuid/chat/ws")
	app.Get("/stream/:ssuid/viewer/ws")
}
