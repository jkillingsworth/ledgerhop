package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/jkillingsworth/ledgerhop/ripple"
)

//go:generate go-bindata -o assets.go -prefix "assets" assets/...

var templates = template.New("ledgerhop")

func registerTemplate(t *template.Template, name string) {
	b := MustAsset(name)
	s := string(b)
	t.New(name).Parse(s)
}

func executeTemplate(w io.Writer, name string, data interface{}) {
	e := templates.ExecuteTemplate(w, name, data)
	if e != nil {
		log.Fatal(e)
	}
}

func handleStyles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprint(w, string(MustAsset("styles.css")))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	executeTemplate(w, "index.html", nil)
}

func handleRipplePing(w http.ResponseWriter, r *http.Request) {

	var req string
	var res string

	format := func(b []byte) string {
		var buffer bytes.Buffer
		e := json.Indent(&buffer, b, "", "  ")
		if e != nil {
			log.Fatal(e)
		}
		return string(buffer.Bytes())
	}

	if r.Method == "POST" {

		r.ParseForm()

		command := struct {
			ID      string `json:"id"`
			Command string `json:"command"`
		}{
			ID:      r.Form.Get("id"),
			Command: r.Form.Get("command"),
		}

		bytes, e := json.Marshal(command)
		if e != nil {
			log.Fatal(e)
		}

		req = format(bytes)
		res = format(ripple.Execute(bytes))
	}

	data := struct {
		Req string
		Res string
	}{
		req,
		res,
	}

	executeTemplate(w, "ripple/ping.html", data)
}

func init() {
	registerTemplate(templates, "index.html")
	registerTemplate(templates, "ripple/ping.html")
}

func main() {
	http.HandleFunc("/styles.css", handleStyles)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/ripple/ping", handleRipplePing)
	http.ListenAndServe(":8080", nil)
}
