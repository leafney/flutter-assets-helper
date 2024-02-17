import { createApp } from 'vue'
import './assets/style.css'
import App from './App.vue'
// import MySoc from './utils/mysoc';
import XSocket from '@/utils/xsocket';

const app = createApp(App)

// const socket = new Websocket('ws://localhost:8080/ws', {
//     rejectUnauthorized: false,
// });

// const soc = new MySoc('ws://localhost:8080/ws');
const soc = new XSocket('ws://localhost:8080/ws');

app.provide('socket', soc);

app.mount('#app')
