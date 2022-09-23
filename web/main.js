let socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
    console.log("Successfully Connected");
    const input = prompt("What's your name?");
    socket.send(input)
};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
    socket.send("Client Closed!")
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
};

socket.onmessage = function(e){
    var server_message = e.data;
    console.log(server_message);
};