package ripple

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func executeCommand(req []byte) (res []byte) {

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

func formatJSON(b []byte) string {
	var buffer bytes.Buffer
	e := json.Indent(&buffer, b, "", "  ")
	if e != nil {
		log.Fatal(e)
	}
	return string(buffer.Bytes())
}

// Handle executes a ripple command.
func Handle(w http.ResponseWriter, r *http.Request, getCommand func(r *http.Request) interface{}) interface{} {

	var req string
	var res string

	if r.Method == "POST" {

		r.ParseForm()
		command := getCommand(r)

		bytes, e := json.Marshal(command)
		if e != nil {
			log.Fatal(e)
		}

		req = formatJSON(bytes)
		res = formatJSON(executeCommand(bytes))
	}

	return struct {
		Req string
		Res string
	}{
		req,
		res,
	}
}

// GetCommandPing creates a "ping" command.
func GetCommandPing(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
	}
}

// GetCommandAccountInfo creates an "account_info" command.
func GetCommandAccountInfo(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
		Ledger  string `json:"ledger_index"`
		Account string `json:"account"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
		Ledger:  r.Form.Get("ledger"),
		Account: r.Form.Get("account"),
	}
}

// GetCommandAccountLines creates an "account_lines" command.
func GetCommandAccountLines(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
		Ledger  string `json:"ledger_index"`
		Account string `json:"account"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
		Ledger:  r.Form.Get("ledger"),
		Account: r.Form.Get("account"),
	}
}
