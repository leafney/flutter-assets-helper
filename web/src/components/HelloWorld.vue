<script setup lang="ts">
// import Websocket from '@/utils/xsocket';

import XSocket from '@/utils/xsocket';
// import MySoc from '@/utils/mysoc';
import { inject, onMounted, ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)

const myWs = inject('socket', new XSocket())
console.log('初始化 myWs', myWs)


onMounted(() => {
  console.log('connected start')
  console.log(myWs)
  myWs.connect();

  myWs.onMessage = (event) => {
    console.log('重写了接收方法', event.data);
  }

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
  <h1>{{ count }}</h1>
  <button @click="fasong">发送</button>
</template>

<style scoped></style>
