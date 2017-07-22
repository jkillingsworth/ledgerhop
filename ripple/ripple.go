package ripple

import (
	"log"

	"github.com/gorilla/websocket"
)

// Execute command
func Execute(req []byte) (res []byte) {

	c, _, e := websocket.DefaultDialer.Dial("wss://s1.ripple.com", nil)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Close()

	e = c.WriteMessage(websocket.TextMessage, req)
	if e != nil {
		log.Fatal(e)
	}

	_, res, e = c.ReadMessage()
	if e != nil {
		log.Fatal(e)
	}

	e = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if e != nil {
		log.Fatal(e)
	}

	return
}
