const BUTTON_ID = 'playersBtn';

function initPlayersButton() {
    const button = document.getElementById(BUTTON_ID);

    button.addEventListener('click', () => {
        document.dispatchEvent(new CustomEvent('show-top-players'));
    });

    return button;
}

export default initPlayersButton;
