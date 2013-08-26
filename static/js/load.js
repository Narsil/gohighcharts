
$(function () {
    var serie;
    function onMessage(e) {
        if (serie){
            var shift;
            if (serie.data.length > 10){
                shift = true;
            }else{
                shift = false;
            }
            serie.addPoint([parseInt(e.data, 10)], true, shift);
        }
    }
    function onClose() {
        console.log("Connection closed.");
    }
        

    $.getJSON('./data/', function(data){
        $('#chart').highcharts(data);
        serie = $('#chart').highcharts().series[0];
    });

    websocket = new WebSocket("ws://"+ window.location.host + window.location.pathname + "streaming/");
    websocket.onmessage = onMessage;
    websocket.onclose = onClose;



});

