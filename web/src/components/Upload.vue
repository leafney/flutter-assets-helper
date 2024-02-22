<script setup lang="ts">
import { protocol } from "@/proto/proto";
import YSocket from "@/utils/ysocket";
import { inject } from "vue";

const myWs = inject('socket', new YSocket(''))

// 上传图片
const customRequest = ({ file }: any) => {
    console.log(file);
    // 直接传输图片文件
    // myWs.sendFile(file.file as File);

    let name = file.name;

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
                fileName: name,
            }

            const message = protocol.Message.create(data);
            myWs.send(protocol.Message.encode(message).finish());
        }
    };
    reader.readAsArrayBuffer(file.file);

}


</script>
<template>
    <div>
        <n-upload directory-dnd :show-file-list="false" :custom-request="customRequest">
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
    </div>
</template>
<style scoped></style>