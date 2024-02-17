import { createApp } from 'vue'
import './assets/style.css'
import App from './App.vue'


const app = createApp(App)

// import MySoc from './utils/mysoc';
// const soc = new MySoc('ws://localhost:8080/ws');

// import XSocket from '@/utils/xsocket';
// const soc = new XSocket('ws://localhost:8080/ws');

// app.provide('socket', soc);

app.mount('#app')
