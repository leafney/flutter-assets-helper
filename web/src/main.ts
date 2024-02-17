import { createApp } from 'vue'
import './assets/style.css'
import App from './App.vue'
import MySoc from './utils/mysoc';
// import Websocket from '@/utils/xsocket';
// import Websocket from './utils/xsocket.js';

const app = createApp(App)

// const socket = new Websocket('ws://localhost:8080/ws', {
//     rejectUnauthorized: false,
// });

const soc = new MySoc('ws://localhost:8080/ws');

app.provide('socket', soc);

app.mount('#app')
