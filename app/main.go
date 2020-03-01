package main

import (
	"angles/converter"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

type OperationType int

const (
	Add OperationType = iota
	Sub
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/doit", handleDoit)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Hiba a szerver indítása közben. %s\n", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatalf("Hiba az index oldal kiszolgálása közben. %s\n", err)
	}
}

func handleDoit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Bejövő adatok feldolgozása sikertelen. %s\n", err)
	}

	opType, err := strconv.Atoi(r.Form.Get("op-type"))
	if err != nil {
		log.Printf("Hibás művelet. %s\n", err)
	}

	valX := r.Form.Get("val-x")
	typeX := r.Form.Get("type-x")
	radX, err := converter.ConvertToRad(valX, typeX)

	valY := r.Form.Get("val-y")
	typeY := r.Form.Get("type-y")
	radY, err := converter.ConvertToRad(valY, typeY)

	var res float64
	switch OperationType(opType) {
	case Add:
		res = radX + radY
	case Sub:
		res = radX - radY
	default:
		// TODO nincs ilyen művelet
	}

	var tplRes interface{}
	resType := r.Form.Get("res-type")
	switch resType {
	case "dms":
		tplRes = converter.ConvertRadToDms(res)
	case "deg":
		tplRes = converter.ConvertRadToDeg(res)
	case "rad":
		tplRes = res
	}

	if err := tpl.ExecuteTemplate(w, "index.html", tplRes); err != nil {
		log.Printf("Hiba az index oldal kiszolgálása közben. %s\n", err)
		// TODO HTTP error
	}
}
