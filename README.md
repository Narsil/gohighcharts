gohighcharts
============

Library to display graphics using highcharts on a local server. Supports dynamic data through channels.


Installing
==========

Within your project run

```bash
// Download the sources
go get github.com/Narsil/gohighcharts

// Get the static files that will be used by the server
git clone https://github.com/Narsil/gohighcharts.git
cp -r gohighcharts/{tmpl,static} .
rm -rf gohighcharts
```

Usage
=====

Simple Chart Example
--------------------


```go
options := map[string]interface{}{
  "series":  []interface{}{
    map[string]interface{}{
      "name": "MyData",
      "data": []int{1, 2, 3},
    },
  },
  "chart": map[string]interface{}{
    "type": "line",
  },
}                                                                          
NewChart("/chart/", options) 
```

Then simply visit `http://localhost:8080/chart/` to see you chart.


Dynamic Chart
-------------

When using a dynamic chart

```go
package main

import (
  highcharts "github.com/Narsil/gohighcharts"
)

func main(){
  data := make(chan interface{})                                             
  options := map[string]interface{}{
      "series":  []interface{}{
          map[string]interface{}{
              "name": "Dynamic chart",
              "data": []int{},
          },
      },
      "chart": map[string]interface{}{
          "type": "line",
      },
  }
  highcharts.NewDynamicChart("/dynamic/", options, data)
  go func(){                                                                 
      for i := 0; i < 10; i++{                                               
          data<-i                                                            
          time.Sleep(i * 1e9)                                                      
      }                                                                      
  }()
}
```

And visit `http://localhost:8080/dyamic/`.
