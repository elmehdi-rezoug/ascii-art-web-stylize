package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
	"ascii-art-web/asciigenerator"
)

type Data struct {
	Output string
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 wrong banner"))
		return
	}

	_, ok := r.Form["input"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 must provide an input"))
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}

	output, err := asciigenerator.GenerateAsciiArt(input, banner)
	if err != nil {
		if strings.HasPrefix(err.Error(), "400") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 internal server error"))
		}
		return
	}

	if strings.HasPrefix(output, "\n") {
		output = "\n" + output
	}

	data := Data{
		Output: output,
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}
	w.Write(buff.Bytes())
}
