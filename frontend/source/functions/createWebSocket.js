const eventNames = {
    open: 'open',
    close: 'close',
    error: 'error',
    message: 'message'
};

const initCallbacks = {
    [eventNames.open]: () => {},
    [eventNames.close]: () => {},
    [eventNames.error]: () => {},
    [eventNames.message]: () => {}
};

function createWebSocket({ url = '', protocols = [], handlerCallbacks = {} }) {
    const callbacks = {
        ...initCallbacks,
        ...handlerCallbacks
    };

    const handlers = {
        [eventNames.open]: () => {
            console.log('Successfully Connected');
            callbacks.open();
        },
        [eventNames.close]: event => {
            console.log('Socket Closed Connection: ', event);
            callbacks.close();
        },
        [eventNames.error]: error => {
            console.log('Socket Error: ', error);
            callbacks.error();
        },
        [eventNames.message]: event => {
            console.log(`server message: ${event.data}`);
            callbacks.message();
        }
    };

    const webSocket = new WebSocket(url, protocols);

    Object.values(eventNames).forEach(event => {
        webSocket.addEventListener(event, handlers[event]);
    });

    return webSocket;
}

export default createWebSocket;
