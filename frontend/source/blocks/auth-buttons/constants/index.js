import { ROOT } from 'source/constants';

const BUTTON_IDS = {
    login: 'login-btn',
    signup: 'signup-btn'
};

const TEMPLATE_PATHS = {
    login: `${ROOT}/templates/login.html`,
    signup: `${ROOT}/templates/signup.html`
};

const FORM_NAMES = {
    login: 'loginForm',
    signup: 'signupForm'
};

export const BUTTON_OPTIONS = ['login', 'signup'].map(button => ({
    id: BUTTON_IDS[button],
    templatePath: TEMPLATE_PATHS[button],
    formName: FORM_NAMES[button]
}));
