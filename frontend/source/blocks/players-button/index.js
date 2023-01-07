const BUTTON_ID = 'playersBtn';

function initPlayersButton() {
    const button = document.getElementById(BUTTON_ID);

    button.addEventListener('click', () => {
        dispatchEvent(new CustomEvent('show-top-players'));
    });
}

export default initPlayersButton;
