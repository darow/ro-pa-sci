function handleWS() {
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
}

const contentElem = document.querySelector('#content')

document.querySelector('#login-btn').addEventListener("click", () => {
    fetch("/login.html")
        .then((response) => response.text())
        .then((data) => {contentElem.innerHTML = data});
})

document.querySelector('#signup-btn').addEventListener("click", () => {
    fetch("/signup.html")
        .then((response) => response.text())
        .then((data) => {
            contentElem.innerHTML = data
            document.forms['myForm'].addEventListener('submit', (event) => {
                event.preventDefault();
                // TODO do something here to show user that form is being submitted
                fetch(event.target.action, {
                    method: 'POST',
                    body: new URLSearchParams(new FormData(event.target)) // event.target is the form
                }).then((resp) => {
                    return resp.json(); // or resp.text() or whatever the server sends
                }).then((body) => {
                    console.log(body)
                }).catch((error) => {
                    console.log(body)
                });
            });
        });
})