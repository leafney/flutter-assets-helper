
class XSocket {
    constructor(url) {
        this.url = url;
        this.reconnectInterval = null;
    }

    connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = this.onOpen.bind(this);
        this.socket.onmessage = this.onMessage.bind(this);
        this.socket.onclose = this.onClose.bind(this);
        this.socket.onerror = this.onError.bind(this);
    }

    onOpen(event) {
        console.log('WebSocket连接已打开');
        if (this.reconnectInterval) {
            clearInterval(this.reconnectInterval); // 清除重连定时器
            this.reconnectInterval = null;
        }
    }

    onMessage(event) {
        console.log('接收到消息:', event.data);
    }

    onClose(event) {
        if (event.wasClean) {
            console.log('WebSocket连接已正常关闭');
        } else {
            console.log('WebSocket连接异常断开，尝试重新连接...');
            this.reconnect();
        }
    }

    onError(error) {
        console.error('WebSocket连接发生错误:', error);
        this.reconnect();
    }

    reconnect() {
        if (!this.reconnectInterval) {
            this.reconnectInterval = setInterval(() => {
                console.log('尝试重新连接...');
                this.connect();
            }, 3000); // 3秒后尝试重新连接
        }
    }

    send(data) {
        this.socket.send(data);
        // send(JSON.stringify(data));
    }

    sendFile(file) {
        const reader = new FileReader();
        reader.onload = (event) => {
            this.socket.send(event.target.result);
        };
        reader.readAsArrayBuffer(file);
    }

    close() {
        this.socket.close();
    }
}

export default XSocket;

/*
// 使用自定义的WebSocket类
const customSocket = new CustomWebSocket('ws://localhost:3000/ws/123');
customSocket.send('Hello, Backend!');

// 发送文件示例
const fileInput = document.getElementById('fileInput');
fileInput.addEventListener('change', (event) => {
  const file = event.target.files[0];
  customSocket.sendFile(file);
});
*/ 