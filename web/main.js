const apiUri = "localhost:8000"

const contentElem = document.querySelector('#content')
const loginBtn = document.querySelector('#login-btn')
const signupBtn = document.querySelector('#signup-btn')
const logoutBtn = document.querySelector('#logout-btn')
const usernameElem = document.querySelector('#username')
const playersBtn = document.querySelector('#playersBtn')
const socketSendBtn = document.querySelector('#socketSendBtn')

let ws = undefined

checkAuth(showPlayersTop)

playersBtn.addEventListener("click", () => {
    showPlayersTop()
})

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
                        checkAuth(showPlayersTop)
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
                        checkAuth(showPlayersTop)
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

function refreshWS(callback) {
    if (ws!==undefined&&ws.readyState === WebSocket.OPEN) {
        ws.close()
    }
    ws = new WebSocket(`ws://${apiUri}/auth/ws`);
    configureWS(callback)
}

function configureWS(callback = () => {}) {
    ws.onopen = () => {
        console.log("Successfully Connected");
        callback()

        document.forms['socketForm'].addEventListener('submit', (event) => {
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
                    checkAuth(showPlayersTop)
                }
            }).catch((error) => {
                console.log(error)
            });
        });
    };

    ws.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        ws.send("Client Closed!")
    };

    ws.onerror = error => {
        console.log("Socket Error: ", error);
    };

    ws.onmessage = function (e) {
        let server_message = e.data;
        console.log(server_message);
    };
}

function showPlayersTop() {
    fetch("/online_users")
        .then((response) => response.json())
        .then((data) => {
            let playersList = ""
            let counter = 1
            for (const [key, user] of Object.entries(data)) {
                let btn = ``
                if (user.is_online) {
                    btn = `<btn id="invite-${user.id}" class="btn btn-sm btn-success btn-block">пригласить✉</btn>`
                }
                playersList += `<li class="m-1">${counter}. ${user.name} ${user.is_online} ${user.score} ${btn}</li>`
                counter++
            }

            contentElem.innerHTML = playersTemplate.formatUnicorn({playerList: `${playersList}`});

        });
}

function checkAuth(callback) {
    fetch("/auth/")
        .then((response) => response.json())
        .then((user) => {
            logoutBtn.style.display = ""
            loginBtn.style.display = "none"
            signupBtn.style.display = "none"
            usernameElem.innerHTML = user.name
            refreshWS(callback)
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



