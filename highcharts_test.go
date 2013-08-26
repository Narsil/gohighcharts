package highcharts

import (
    "testing"
    "net/http"
    "io/ioutil"
    "time"
    "code.google.com/p/go.net/websocket"
    "strconv"
)

func TestNewChart(t *testing.T){
    options := map[string]interface{}{
        "series":  []interface{}{
            map[string]interface{}{
                "name": "toto",
                "data": []int{1, 2, 3},
            },
        },
        "chart": map[string]interface{}{
            "type": "line",
        },
    }
    NewChart("/chart/", options)

    resp, err := http.Get("http://localhost:8080/chart/")
    if err != nil{
        t.Errorf("Error while loading chart page", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        t.Errorf("Error while reading chart page", err)
    }
    // fmt.Println(string(body))

    resp, err = http.Get("http://localhost:8080/chart/data/")
    if err != nil{
        t.Errorf("Error while loading chart data page", err)
    }
    defer resp.Body.Close()
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil{
        t.Errorf("Error while reading chart data page", err)
    }
    if string(body) != "{\"chart\":{\"type\":\"line\"},\"series\":[{\"data\":[1,2,3],\"name\":\"toto\"}]}"{
        t.Errorf("Error in data json")
    }


    resp, err = http.Get("http://localhost:8080/static/js/highcharts.js")
    if err != nil{
        t.Errorf("Error while reading static page", err)
    }

    resp, err = http.Get("http://localhost:8080/static/js/load.js")
    if err != nil{
        t.Errorf("Error while reading static page", err)
    }
}

func TestDynamicChart(t *testing.T){
    options := map[string]interface{}{
        "series":  []interface{}{
            map[string]interface{}{
                "name": "toto",
                "data": []int{1, 2, 3},
            },
        },
    }

    data := make(chan interface{})
    NewDynamicChart("/dynamic/", options, data)
    go func(){
        for i := 0; i < 10; i++{
            data<-i
            // Uncomment to test in browser
            // time.Sleep(1e9)
            time.Sleep(1)
        }
    }()
    resp, err := http.Get("http://localhost:8080/dynamic/")
    if err != nil{
        t.Errorf("Error while loading dynamic chart data page", err)
    }
    defer resp.Body.Close()
    _, err = ioutil.ReadAll(resp.Body)
    if err != nil{
        t.Errorf("Error while reading dynamic data page", err)
    }

    origin := "http://localhost/"
    url := "ws://localhost:8080/dynamic/streaming/"
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        t.Errorf("Cannot open websocket", err)
    }
    var msg = make([]byte, 512)
    var n int
    for i := 0; i< 10; i++{
        n, err = ws.Read(msg)
        if err != nil {
            t.Errorf("Cannot receive on websocket", err)
        }
        strmsg := string(msg[:n])
        expmsg := strconv.Itoa(i)
        if strmsg != expmsg {
            t.Errorf("Did not receive the correct message on websocket exepected ", expmsg, " got ", strmsg)
        }
    }

    // Uncomment to test in browser
    // time.Sleep(1e11)
}
