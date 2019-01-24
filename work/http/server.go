package main

import (
    "html/template"
    "io/ioutil"
    "net/http"
)

type Page struct {
    Title string
    Body []byte
}

//MAIN
func main() {
    http.HandleFunc("/index/", index_handler)
    http.HandleFunc("/ver_boletos/", ver_boletos_handler)
    http.ListenAndServe(":8080", nil)
}

//UTILS
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// HANDLERS
func index_handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("/home/ec2-user/work/http/templates/index.html")
    t.Execute(w, nil)
}
func ver_boletos_handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("/home/ec2-user/work/http/templates/ver_boletos.html")
    t.Execute(w, nil)
}
