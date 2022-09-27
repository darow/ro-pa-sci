function handleWS() {
    let socket = new WebSocket("ws://localhost:8080/rps/ws");

    socket.onopen = () => {
        console.log("Successfully Connected");
        socket.send("playersTop")
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
        .then((data) => {
            contentElem.innerHTML = data
            document.forms['loginForm'].addEventListener('submit', (event) => {
                event.preventDefault();
                // TODO do something here to show user that form is being submitted
                fetch(event.target.action, {
                    method: 'POST',
                    body: new URLSearchParams(new FormData(event.target)) // event.target is the form
                }).then((resp) => {
                    return resp.json();
                }).then((body) => {
                    if (body.error) {
                        document.querySelector('#error').textContent = body.error
                    } else {
                        handleWS()
                    }
                }).catch((error) => {
                    console.log(error)
                });
            });
        });
})

document.querySelector('#signup-btn').addEventListener("click", () => {
    fetch("/signup.html")
        .then((response) => response.text())
        .then((data) => {
            contentElem.innerHTML = data
            document.forms['signupForm'].addEventListener('submit', (event) => {
                event.preventDefault();
                // TODO do something here to show user that form is being submitted
                fetch(event.target.action, {
                    method: 'POST',
                    body: new URLSearchParams(new FormData(event.target)) // event.target is the form
                }).then((resp) => {
                    return resp.json();
                }).then((body) => {
                    if (body.error) {
                        document.querySelector('#error').textContent = body.error
                    } else {
                        document.querySelector('#login-btn').click()
                    }
                    }).catch((error) => {
                    console.log(error)
                });
            });
        });
})