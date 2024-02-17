<script setup lang="ts">

import { onMounted, ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)
const connectSts = ref(0)

// import MySoc from '@/utils/mysoc';

// import XSocket from '@/utils/xsocket';
// const myWs = inject('socket', new XSocket())
// console.log('初始化 myWs', myWs)

import YSocket from '@/utils/ysocket';

const myWs = new YSocket('ws://localhost:8080/ws');

myWs.openCallback = (ev: Event) => {
  console.log('11openCallback ', ev);
}

myWs.messageCallback = (ev: MessageEvent) => {
  console.log('11messageCallback ', ev.data, myWs.status);
}

myWs.closeCallback = (ev: any) => {
  console.log('11closeCallback ', ev);
}

myWs.statusChangedCallback = (s: number) => {
  console.log('statusChangedCallback ', s);
  connectSts.value = s;
}


onMounted(() => {
  console.log('connected start')
  console.log(myWs)
  myWs.connect();


  console.log('connected end');
})

const fasong = () => {
  // myWs.send('hello')

  myWs.send('okok heha');
  console.log('点击了按钮');

}

</script>

<template>
  <h1>{{ msg }}</h1>
  <h1>{{ count }}--连接状态：{{ connectSts }}</h1>
  <button @click="fasong">发送</button>
</template>

<style scoped></style>
