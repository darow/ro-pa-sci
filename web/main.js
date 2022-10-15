const contentElem = document.querySelector('#content')
const loginBtn = document.querySelector('#login-btn')
const signupBtn = document.querySelector('#signup-btn')
const logoutBtn = document.querySelector('#logout-btn')
const usernameElem = document.querySelector('#username')
let webSocket = undefined

showPlayersTop()
checkAuth()

loginBtn.addEventListener("click", () => {
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
                        checkAuth()
                    }
                }).catch((error) => {
                    console.log(error)
                });
            });
        });
})

signupBtn.addEventListener("click", () => {
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

logoutBtn.addEventListener("click", () => {
    fetch("/auth/logout")
        .then((response) => {
                if (response.status == 200) {
                    checkAuth()
                }

                showPlayersTop()
            }
        )
})

function handleWS() {
    webSocket = new WebSocket("ws://localhost:8080/auth/ws");

    socket.onopen = () => {
        console.log("Successfully Connected");
        showPlayersTop()
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = function (e) {
        var server_message = e.data;
        console.log(server_message);
    };
}

function showPlayersTop() {
    fetch("/online_users")
        .then((response) => response.json())
        .then((data) => {
            let playersList = ""
            let counter = 1
            for (const [key, value] of Object.entries(data)) {
                playersList += `<li>${counter}. ${value.login} ${value.is_online} ${value.score}</li>`
                counter++
            }

            contentElem.innerHTML = playersTemplate.formatUnicorn({playerList: `${playersList}`});

        });
}

function checkAuth() {
    fetch("/auth/")
        .then((response) => response.json())
        .then((user) => {
            logoutBtn.style.display = ""
            loginBtn.style.display = "none"
            signupBtn.style.display = "none"
            usernameElem.innerHTML = user.login
            handleWS()
        }).catch((error) => {
        logoutBtn.style.display = "none"
        loginBtn.style.display = ""
        signupBtn.style.display = ""
        usernameElem.innerHTML = ""
        // if (webSocket!==undefined&&webSocket.readyState === WebSocket.OPEN) {
        //     webSocket.close()
        // }
    });
}



