gohighcharts
============

Library to display graphics using highcharts on a local server. Supports dynamic data through channels.

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
  highcharts.NewDynamicChart("/dynamic/", nil, data)
  go func(){                                                                 
      for i := 0; i < 10; i++{                                               
          data<-i                                                            
          time.Sleep(i * 1e9)                                                      
      }                                                                      
  }()
}
```

And visit `http://localhost:8080/dyamic/`.
