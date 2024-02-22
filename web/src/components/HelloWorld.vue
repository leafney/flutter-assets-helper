<script setup lang="ts">

import { onMounted, ref } from 'vue'

import ColorValue from '@/components/ColorValue.vue';

// import protobuf from "protobufjs";

import { protocol } from "@/proto/proto";

defineProps<{ msg: string }>()

const count = ref(0)
const connectSts = ref(0)

// import MySoc from '@/utils/mysoc';

// import XSocket from '@/utils/xsocket';
// const myWs = inject('socket', new XSocket())
// console.log('初始化 myWs', myWs)

import YSocket from '@/utils/ysocket';

const myWs = new YSocket('ws://localhost:8080/ws', 30000);

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

// 发送文字
const fasong = () => {

  // myWs.send('okok heha');

  let data = {
    type: 2,
    content: "okok hahaha",
    contentType: 1,
    fileSuffix: "",
  }

  const message = protocol.Message.create(data);
  myWs.send(protocol.Message.encode(message).finish());

  console.log('点击了按钮');

}

// 上传图片
const customRequest = ({ file }: any) => {
  console.log(file);
  // 直接传输图片文件
  // myWs.sendFile(file.file as File);

  // let name = file.name;

  // 使用protobuf方式封装后传输
  let reader = new FileReader();
  reader.onload = (event: ProgressEvent<FileReader>) => {
    if (event.target && event.target.result) {
      let theFile = new Uint8Array(event.target.result as ArrayBuffer);
      let data = {
        type: 2,
        content: "",
        contentType: 3,
        file: theFile,
        fileSuffix: ".png",
      }

      const message = protocol.Message.create(data);
      myWs.send(protocol.Message.encode(message).finish());
    }
  };
  reader.readAsArrayBuffer(file.file);

}

</script>

<template>
  <h1>{{ msg }}</h1>
  <h1>{{ count }}--连接状态：{{ connectSts }}</h1>
  <button @click="fasong">发送</button>

  <n-upload directory-dnd :custom-request="customRequest">
    <n-upload-dragger>
      <div style="margin-bottom: 12px">
        <n-icon size="48" :depth="3">
          <archive-icon />
        </n-icon>
      </div>
      <n-text style="font-size: 16px">
        点击或者拖动文件到该区域来上传
      </n-text>
      <n-p depth="3" style="margin: 8px 0 0 0">
        请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
      </n-p>
    </n-upload-dragger>
  </n-upload>

  <br>
  <ColorValue></ColorValue>
</template>


<style scoped></style>
