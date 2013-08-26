package highcharts

import (
    "net/http"
    "html/template"
    "log"
    "encoding/json"
    "fmt"
    "code.google.com/p/go.net/websocket"
)

var addedStatic = false

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    t := template.New("base")
    t, err := template.ParseFiles("tmpl/base.html")
    if err != nil{
        log.Fatal(err)
    }
    t.Execute(w, nil)
}

func NewChart(url string, options interface{}){
    http.HandleFunc(url, IndexHandler)
    http.HandleFunc(url + "data/", func(w http.ResponseWriter, r *http.Request) {
        opts, err := json.Marshal(options)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Fprintf(w, string(opts))
    })
    if (!addedStatic){
        http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
            http.ServeFile(w, r, r.URL.Path[1:])
        })
        addedStatic = true
    }

    go http.ListenAndServe(":8080", nil)
    return
}

func NewDynamicChart(url string, options interface{}, channel chan interface{}){
    http.HandleFunc(url, IndexHandler)
    http.HandleFunc(url + "data/", func(w http.ResponseWriter, r *http.Request) {
        opts, err := json.Marshal(options)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Fprintf(w, string(opts))
    })
    http.Handle(url + "streaming/", websocket.Handler(func (c *websocket.Conn){
        for ;; {
            data := <-channel
            fmt.Fprint(c, data)
        }
    }))
    if (!addedStatic){
        http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
            http.ServeFile(w, r, r.URL.Path[1:])
        })
        addedStatic = true
    }

    go http.ListenAndServe(":8080", nil)
    return
}
