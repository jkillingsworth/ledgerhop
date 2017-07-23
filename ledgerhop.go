package main

import (
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
	} else {
		executeTemplate(w, "index.html", nil)
	}
}

func handleRipplePing(w http.ResponseWriter, r *http.Request) {
	name := "ripple/ping.html"
	data := ripple.Handle(w, r, ripple.GetCommandPing)
	executeTemplate(w, name, data)
}

func handleRippleAccountInfo(w http.ResponseWriter, r *http.Request) {
	name := "ripple/account-info.html"
	data := ripple.Handle(w, r, ripple.GetCommandAccountInfo)
	executeTemplate(w, name, data)
}

func handleRippleAccountLines(w http.ResponseWriter, r *http.Request) {
	name := "ripple/account-lines.html"
	data := ripple.Handle(w, r, ripple.GetCommandAccountLines)
	executeTemplate(w, name, data)
}

func init() {
	registerTemplate(templates, "index.html")
	registerTemplate(templates, "ripple/ping.html")
	registerTemplate(templates, "ripple/account-info.html")
	registerTemplate(templates, "ripple/account-lines.html")
}

func main() {
	http.HandleFunc("/styles.css", handleStyles)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/ripple/ping", handleRipplePing)
	http.HandleFunc("/ripple/account-info", handleRippleAccountInfo)
	http.HandleFunc("/ripple/account-lines", handleRippleAccountLines)
	http.ListenAndServe(":8080", nil)
}
