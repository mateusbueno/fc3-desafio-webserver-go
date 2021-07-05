package main

import (
	"html/template"
    "net/http"
)

type SimplePageData struct {
    PageTitle string
    ImgSrc    string
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := SimplePageData{
            PageTitle: "Imersão Full Cycle",
            ImgSrc: "https://fw.atarde.uol.com.br/2018/12/750_filme-vingadores-longa-trailer-divulgacao_20181271441192.jpg",
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8000", nil)
}
