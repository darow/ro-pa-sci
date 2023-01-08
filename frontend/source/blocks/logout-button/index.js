import { BUTTON_ID, API_LOGOUT } from '@/logout-button/constants';

function initLogoutButton() {
    const button = document.getElementById(BUTTON_ID);

    button.addEventListener('click', async () => {
        const response = await fetch(API_LOGOUT);

        if (response.status === 200) {
            document.dispatchEvent(new CustomEvent('update-authorization'));
        }

        document.dispatchEvent(new CustomEvent('show-top-players'));
    });

    return button;
}

export default initLogoutButton;
