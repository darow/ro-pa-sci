import { BUTTON_OPTIONS } from '@/auth-buttons/constants';

function initAuthButtons() {
    BUTTON_OPTIONS.forEach(button => {
        const element = document.getElementById(button.id);

        element.addEventListener('click', async () => {
            try {
                const response = await fetch(button.templatePath);
                const text = await response.text();
                const detail = { text, formName: button.formName };

                document.dispatchEvent(new CustomEvent('update-content', { detail }));
            } catch (error) {
                console.log(error);
            }
        });
    });

    const buttons = BUTTON_OPTIONS.map(button => document.getElementById(button.id));

    return buttons;
}

export default initAuthButtons;
