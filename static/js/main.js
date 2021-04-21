
var host = "localhost:8080"
var nombre = "Roberto"
var ws = null
bnt = document.getElementById('bntSocket')
bnt.onclick = function () {
    nombre = document.getElementById('nombre').value
    ws = new WebSocket(`ws://${host}/websocket?nombre=${nombre}`)
    ws.onopen = function() {
    console.log('Conexión establecida')
    }

    ws.onmessage = function (event) {
        console.log(event.data)
    }
}


function SendMsg(msg) {
    var msg = {
        type: "message",
        text: msg,
        Nombre: nombre,
        date: Date.now()
    }
    ws.send(JSON.stringify(msg))
    console.log('Msg enviado')
}


var dato = document.getElementById('dato')
dato.innerText = "Hola, data"
