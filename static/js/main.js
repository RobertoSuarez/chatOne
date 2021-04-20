
var host = "localhost:8080"
var nombre = "Roberto"
var ws = new WebSocket(`ws://${host}/websocket?nombre=${nombre}`)

ws.onopen = function() {
    console.log('Conexión establecida')
}

ws.onmessage = function (event) {
    console.log(event.data)
}

function SendMsg(msg) {
    var msg = {
        type: "message",
        text: msg,
        date: Date.now()
    }
    ws.send(JSON.stringify(msg))
    console.log('Msg enviado')
}


var dato = document.getElementById('dato')
dato.innerText = "Hola, data"
