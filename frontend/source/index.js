import { API, API_URI } from 'frontend/source/constants';
import {
    createWebSocket,
    getTopPlayers
} from 'frontend/source/functions';
import {
    initAuthButtons,
    initLogoutButton,
    initPlayersButton
} from '@';

const authButtons = initAuthButtons();
const logoutButton = initLogoutButton();

initPlayersButton();

const contentElem = document.getElementById('content');
const errorElem = document.getElementById('error');
const usernameElem = document.getElementById('username');

let ws;

function closeWebSocket() {
    if (!ws) {
        return;
    }

    ws.close();
}

function refreshWS() {
    closeWebSocket();
    ws = createWebSocket({
        url: `ws://${API_URI}${API.authWs}`,
        protocols: {},
        handlerCallbacks: { open: updateTopPlayers }
    });
}

async function updateTopPlayers() {
    try {
        const response = await fetch(API.onlineUsers);
        const players = await response.json();

        contentElem.innerHTML = getTopPlayers(players);
    } catch (error) {
        console.log(error);
    }
}

async function handleSubmitForm(event) {
    event.preventDefault();

    try {
        const target = event.target;
        const response = await fetch(target.action, {
            method: 'POST',
            body: new URLSearchParams(new FormData(target))
        });
        const data = await response.json();

        if (data.error) {
            const detail = { error: data.error };

            document.dispatchEvent(new CustomEvent('submit-error', { detail }));

            return;
        }

        document.dispatchEvent(new CustomEvent('update-authorization'));
    } catch (error) {
        console.log(error);
    }
}

document.addEventListener('submit-error', ({ detail }) => {
    errorElem.textContent = detail.error;
});

document.addEventListener('show-top-players', updateTopPlayers);

document.addEventListener('update-content', ({ detail }) => {
    contentElem.innerHTML = detail.text;
    document.forms[detail.formName].addEventListener('submit', handleSubmitForm);
});

document.addEventListener('update-authorization', async () => {
    try {
        const response = await fetch(API.auth);
        const user = await response.json();

        usernameElem.innerHTML = user.name;
        logoutButton.hidden = false;
        authButtons.forEach(button => {
            button.hidden = true;
        });

        refreshWS();
    } catch (error) {
        usernameElem.innerHTML = '';
        logoutButton.hidden = true;
        authButtons.forEach(button => {
            button.hidden = false;
        });
    }
});
