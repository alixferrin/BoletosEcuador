package main

import (
    "html/template"
    "io/ioutil"
    "net/http"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Page struct {
    Title string
    Body []byte
}

type Evento struct {
    Id int
    Nombre string
    Id_venue int
    Fecha string
}

//MAIN
func main() {
    http.HandleFunc("/index/", index_handler)
    http.HandleFunc("/ver_boletos/", ver_boletos_handler)
    http.HandleFunc("/evento/", ver_evento_handler)
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
    eventos := getEventos("SELECT idEvento, nombre, idVenue, fecha  FROM Evento")
    log.Println(eventos)
    t, _ := template.ParseFiles("/home/ec2-user/work/http/templates/ver_boletos.html")
    t.Execute(w, eventos)
}
func ver_evento_handler(w http.ResponseWriter, r *http.Request) {
    id_evento := r.URL.Path[len("/evento/"):]
    evento := getEvento("SELECT idEvento, nombre, idVenue, fecha  FROM Evento WHERE idEvento=?", id_evento)
    log.Println(evento)
    t, _ := template.ParseFiles("/home/ec2-user/work/http/templates/ver_evento.html")
    t.Execute(w, evento)
}


//SQL
func getEventos(query string) *[]Evento {
    db, err := sql.Open("mysql", "remoteu:password@tcp(ec2-35-182-205-119.ca-central-1.compute.amazonaws.com:3306)/BoletosEcuador")
    if err != nil {
        panic(err)
    }

    rows, err := db.Query(query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    eventos := []Evento{}
    for rows.Next() {
        evento := Evento{}
        err = rows.Scan(&evento.Id, &evento.Nombre, &evento.Id_venue, &evento.Fecha)
        if err != nil {
            panic(err)
        }
        eventos = append(eventos, evento)
    }

    err = rows.Err()
    if err != nil {
        panic(err)
    }

    return &eventos
}
func getEvento(query string, id_evento string) *Evento {
    db, err := sql.Open("mysql", "remoteu:password@tcp(ec2-35-182-205-119.ca-central-1.compute.amazonaws.com:3306)/BoletosEcuador")
    if err != nil {
        panic(err)
    }

    evento := Evento{}
    row := db.QueryRow(query, id_evento)
    err = row.Scan(&evento.Id, &evento.Nombre, &evento.Id_venue, &evento.Fecha)

    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("Zero rows found")
        } else {
            panic(err)
        }
    }

    return &evento
}
