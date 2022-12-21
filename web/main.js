const apiUri = "localhost:8000"

const contentElem = document.querySelector('#content')
const loginBtn = document.querySelector('#login-btn')
const signupBtn = document.querySelector('#signup-btn')
const logoutBtn = document.querySelector('#logout-btn')
const usernameElem = document.querySelector('#username')
const playersBtn = document.querySelector('#playersBtn')
let ws

checkAuth(showPlayersTop)

playersBtn.addEventListener('click', () => {
    showPlayersTop()
})

loginBtn.addEventListener('click', async () => {
    try {
        await fetchTemplate('/login.html', 'loginForm')
    } catch (error) {
        console.log(error)
    }
})
signupBtn.addEventListener('click', async () => {
    try {
        await fetchTemplate('/signup.html', 'signupForm')
    } catch (error) {
        console.log(error)
    }
})

async function fetchTemplate(template, formName) {
    const response = await fetch(template)

    contentElem.innerHTML = await response.text()
    document.forms[formName].addEventListener('submit', handleSubmitForm)
}

async function handleSubmitForm(event) {
    event.preventDefault();

    // TODO do something here to show user that form is being submitted
    const response = await fetch(event.target.action, {
        method: 'POST',
        body: new URLSearchParams(new FormData(event.target)) // event.target is the form
    })
    const data = await response.json()

    if (data.error) {
        document.querySelector('#error').textContent = data.error
        return
    }

    checkAuth(showPlayersTop)
}

logoutBtn.addEventListener('click', async () => {
    const response = await fetch('/auth/logout')

    if (response.status === 200) {
        await checkAuth()
    }

    await showPlayersTop()
})

function refreshWS(callback) {
    closeWebSocket(ws)
    ws = createWebSocket(`ws://${apiUri}/auth/ws`, { open: callback })
}

function createWebSocket(url, options, handlerCallbacks) {
    if (!url) {
        console.log('Не передан URL!')
    }

    const eventNames = {
        open: 'open',
        close: 'close',
        error: 'error',
        message: 'message',
    }

    const callbacks = {
        [eventNames.open]: () => {},
        [eventNames.close]: () => {},
        [eventNames.error]: () => {},
        [eventNames.message]: () => {},
        ...handlerCallbacks
    }

    const handlers = {
         [eventNames.open]: () => {
            console.log('Successfully Connected')
            callbacks.open()
        },
        [eventNames.close]: event => {
            console.log('Socket Closed Connection: ', event);
            callbacks.close()
        },
        [eventNames.error]: error => {
            console.log('Socket Error: ', error);
            callbacks.error()
        },
        [eventNames.message]: () => {
            console.log(`server message: ${e.data}`)
            callbacks.message()
        },
    }

    const webSocket = new WebSocket(url, options)

    Object.values(eventNames).forEach(event => {
        webSocket.addEventListener(event, handlers[event])
    })

    return webSocket
}

function closeWebSocket(webSocket) {
    if (!webSocket) {
        return
    }

    webSocket.close()
}

async function showPlayersTop() {
    const response = await fetch('/online_users')
    const data = await response.json()

    const playersList = Object.values(data).map((user, index) => {
        return `<li class="m-1">${index + 1}. ${getUserInfoElement(user)}</li>`
    })

    contentElem.innerHTML = playersTemplate.formatUnicorn({ playerList: playersList });
}

function getUserInfoElement({ name, is_online, score, id }) {
    const inviteButton = `<btn id="invite-${id}" class="btn btn-sm btn-success btn-block">пригласить✉</btn>`

    return `${name} ${is_online} ${score} ${inviteButton}`
}

async function checkAuth(callback) {
    try {
        const response = await fetch('/auth/')
        const user = await response.json()

        logoutBtn.hidden = false
        loginBtn.hidden = true
        signupBtn.hidden = true
        usernameElem.innerHTML = user.name
        refreshWS(callback)
    } catch (error) {
        logoutBtn.hidden = true
        loginBtn.hidden = false
        signupBtn.hidden = false
        usernameElem.innerHTML = ''
        // if (webSocket!==undefined&&webSocket.readyState === WebSocket.OPEN) {
        //     webSocket.close()
        // }
    }
}
