package highcharts

import (
    "net/http"
    "html/template"
    "log"
    "encoding/json"
    "fmt"
    "code.google.com/p/go.net/websocket"
)

var started = false
var static = false
var port = ":8080"

// Creates a new highchart based on options argument at url
// View at http://localhost:8080/url/
// To change port see SetPort
func NewChart(url string, options interface{}){
    http.HandleFunc(url, indexHandler)
    http.HandleFunc(url + "data/", func(w http.ResponseWriter, r *http.Request) {
        opts, err := json.Marshal(options)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Fprintf(w, string(opts))
    })

    eventualServerStart()
}

// Creates a new highchart based on options argument at url
// Whenever you add data to your channel it will be sent to
// your graph via websocket.
// Only support one channel on one graph for now.
// View at http://localhost:8080/url/
// To change port see SetPort
func NewDynamicChart(url string, options interface{}, channel chan interface{}){
    http.HandleFunc(url, indexHandler)
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
    eventualServerStart()

}

// Sets the port for the server to see the graphs
// Is in the same format as ListenAndServe, ":8080"
func SetPort(server_port string){
    port = server_port
    started = false
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t := template.New("base")
    t, err := template.ParseFiles("tmpl/base.html")
    if err != nil{
        log.Fatal(err)
    }
    t.Execute(w, nil)
}

func eventualServerStart(){
    if !started{
        if !static {
            http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, r.URL.Path[1:])
            })
            static = true
        }
        go func(){
            started = true
            log.Fatal(http.ListenAndServe(port, nil))
        }()
    }
}
