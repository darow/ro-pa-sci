import { BUTTON_OPTIONS } from '@/auth-buttons/constants';

function initAuthButtons() {
    BUTTON_OPTIONS.forEach(button => {
        const element = document.getElementById(button.id);

        element.addEventListener('click', async () => {
            try {
                const response = await fetch(button.templatePath);
                const text = await response.text();
                const detail = { text };

                dispatchEvent(new CustomEvent('update-content', { detail }));
            } catch (error) {
                console.log(error);
            }
        });

        document.forms[button.formName].addEventListener('submit', handleSubmitForm);
    });
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

            dispatchEvent(new CustomEvent('submit-error', { detail }));

            return;
        }

        dispatchEvent(new CustomEvent('update-authorization'));
    } catch (error) {
        console.log(error);
    }
}

export default initAuthButtons;
