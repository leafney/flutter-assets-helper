
class MySoc {
    constructor(url) {
        this.url = url;
        this.ws = null;
    }

    connect() {
        this.ws = new WebSocket(this.url);

        this.ws.onopen = () => {
            console.log('websocket connection opened.');
        }

        this.ws.onmessage = (event) => {
            console.log('websocket message received.', event.data);
        }

    }

    send(data) {
        if (this.ws.readyState === WebSocket.OPEN) {
            this.ws.send(data);
        }
    }

    close() {
        this.ws.close();
    }
}

export default MySoc;