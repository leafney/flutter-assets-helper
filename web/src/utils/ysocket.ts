
class YSocket {
    url: string;
    private socket!: WebSocket;
    private reconnectInterval: NodeJS.Timeout | null;
    status: number; // 0: 未连接, 1: 连接中, 2: 已连接, 3: 已关闭, 4: 断线重连, 5: 连接异常
    retryTime: number;
    openCallback?: (event: Event) => void;
    messageCallback?: (event: MessageEvent) => void;
    closeCallback?: (event: CloseEvent) => void;
    errorCallback?: (error: Event) => void;
    statusChangedCallback?: (status: number) => void;

    constructor(
        url: string,
        retryTime: number = 5000,// 设置默认的重连时间
        openCallback?: (event: Event) => void,
        messageCallback?: (event: MessageEvent) => void,
        closeCallback?: (event: CloseEvent) => void,
        errorCallback?: (error: Event) => void,
        statusChangedCallback?: (status: number) => void,
    ) {
        this.reconnectInterval = null;
        this.status = 0; // 状态，未连接
        this.url = url;
        this.openCallback = openCallback;
        this.messageCallback = messageCallback;
        this.closeCallback = closeCallback;
        this.errorCallback = errorCallback;
        this.statusChangedCallback = statusChangedCallback;
        this.retryTime = retryTime;
    }

    connect(): void {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = this.onOpen.bind(this);
        this.socket.onmessage = this.onMessage.bind(this);
        this.socket.onclose = this.onClose.bind(this);
        this.socket.onerror = this.onError.bind(this);

        this.setStatus(1); // 连接中
    }

    close(): void {
        if (this.socket) {
            // 已关闭
            this.setStatus(3);
            this.socket.close();
        }
    }

    onOpen(event: Event): void {
        console.log('WebSocket连接已打开');
        // 已连接
        this.setStatus(2);
        if (this.reconnectInterval) {
            clearInterval(this.reconnectInterval); // 清除重连定时器
            this.reconnectInterval = null;
        }
        // 回调函数
        if (this.openCallback) {
            this.openCallback(event);
        }
    }

    onMessage(event: MessageEvent): void {
        console.log('接收到消息:', event.data);
        if (this.messageCallback) {
            this.messageCallback(event);
        }
    }

    onClose(event: CloseEvent): void {
        // 已关闭
        this.setStatus(3);
        if (event.wasClean) {
            console.log('WebSocket连接已正常关闭');
            if (this.closeCallback) {
                this.closeCallback(event);
            }
        } else {
            console.log('WebSocket连接异常断开，尝试重新连接...');
            this.reconnect();
        }
    }

    onError(error: Event): void {
        // 连接异常
        this.setStatus(5);
        console.error('WebSocket连接发生错误:', error);
        this.reconnect();
        // 回调函数
        if (this.errorCallback) {
            this.errorCallback(error);
        }
    }

    reconnect(): void {
        // 断线重连
        this.setStatus(4);
        if (!this.reconnectInterval) {
            this.reconnectInterval = setInterval(() => {
                console.log('尝试重新连接...');
                this.connect();
            }, this.retryTime); // 3秒后尝试重新连接
        }
    }

    send(data: string | ArrayBufferLike | Blob | ArrayBufferView): void {
        this.socket.send(data);
    }

    sendFile(file: File): void {
        const reader = new FileReader();
        reader.onload = (event: ProgressEvent<FileReader>) => {
            if (event.target && event.target.result) {
                this.socket.send(event.target.result as ArrayBuffer);
            }
        };
        reader.readAsArrayBuffer(file);
    }

    setStatus(status: number) {
        this.status = status;
        if (this.statusChangedCallback) {
            this.statusChangedCallback(status);
        }
    }
}

export default YSocket;

/* 第一版实现
class YSocket {
    constructor(url, openCallback, messageCallback, errorCallback) {
        this.url = url;
        this.reconnectInterval = null;
        this.status = 0; // 状态，未连接
        this.openCallback = openCallback;
        this.messageCallback = messageCallback;
        this.errorCallback = errorCallback;
    }

    connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = this.onOpen.bind(this);
        this.socket.onmessage = this.onMessage.bind(this);
        this.socket.onclose = this.onClose.bind(this);
        this.socket.onerror = this.onError.bind(this);
    }

    close() {
        if (this.socket) {
            this.status = 2; // 已关闭
            this.socket.close();
        }
    }

    onOpen(event) {
        console.log('WebSocket连接已打开');
        this.status = 1; // 已连接
        if (this.reconnectInterval) {
            clearInterval(this.reconnectInterval); // 清除重连定时器
            this.reconnectInterval = null;
        }
        // 回调函数
        if (this.openCallback) {
            this.openCallback(event);
        }
    }

    onMessage(event) {
        console.log('接收到消息:', event.data);
        if (this.messageCallback) {
            this.messageCallback(event);
        }
    }

    onClose(event) {
        this.status = 2; // 已关闭
        if (event.wasClean) {
            console.log('WebSocket连接已正常关闭');
        } else {
            console.log('WebSocket连接异常断开，尝试重新连接...');
            this.reconnect();
        }
    }

    onError(error) {
        this.status = 4; // 连接异常
        console.error('WebSocket连接发生错误:', error);
        this.reconnect();
        // 回调函数
        if (this.errorCallback) {
            this.errorCallback(error);
        }
    }

    reconnect() {
        this.status = 3; // 断线重连
        if (!this.reconnectInterval) {
            this.reconnectInterval = setInterval(() => {
                console.log('尝试重新连接...');
                this.connect();
            }, 3000); // 3秒后尝试重新连接
        }
    }

    send(data) {
        this.socket.send(data);
    }

    sendFile(file) {
        const reader = new FileReader();
        reader.onload = (event) => {
            this.socket.send(event.target.result);
        };
        reader.readAsArrayBuffer(file);
    }

    status() {
        return this.status;
    }
}

export default YSocket;
*/


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