// import WebSocket from 'ws';
import WebSocket from 'isomorphic-ws';

class Websocket {

    constructor(url, options = {}) {
        this.url = url;
        this.options = options;
        this.ws = null;
    }

    connect() {
        this.ws = new WebSocket(this.url, this.options);

        this.ws.onopen = () => {
            console.log('websocket connection opened.');
        };
        this.ws.onmessage = (event) => {
            console.log('websocket message received.', event.data);
        };
        this.ws.onerror = (error) => {
            console.error('websocket error occurred.', error);
        };
        this.ws.onclose = () => {
            console.log('websocket connection closed');
        };

    }

    send(data) {
        if (this.ws.readyState === WebSocket.OPEN) {
            this.ws.send(data);
        } else {
            console.error('websocket connection not open.');
        }
    }

    close() {
        this.ws.close();
    }
}
export default Websocket;