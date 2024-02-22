<script setup lang="ts">
import { protocol } from "@/proto/proto";
import YSocket from "@/utils/ysocket";
import { inject, ref } from "vue";

const colorValue = ref('')
// #FF0F1419
const colorList = ref<string[]>([]);

const myWs = inject('socket', new YSocket(''));

const copyColor = async () => {
    try {
        let copyText = await navigator.clipboard.readText();
        if (copyText) {
            colorValue.value = copyText;
            colorList.value.push(copyText);

            let data = {
                type: 2,
                content: copyText,
                contentType: 2,
            }

            const message = protocol.Message.create(data);
            myWs.send(protocol.Message.encode(message).finish());

        } else {
            console.log('剪切板为空');
        }
    } catch (error) {
        console.error(error);
    }
}

const handleClose = (index: number) => {
    // colorList.value.slice(index, 1);
    console.log(index);
}

</script>
<template>
    <div>
        <button @click="copyColor">粘贴色值</button>
        <h3>{{ colorValue }}</h3>

        <div>
            <ul>
                <li v-for="item, index in colorList">
                    <n-tag :color="{ color: item, textColor: item, borderColor: '#555' }" closable
                        @close="handleClose(index)">
                        超人不会飞
                    </n-tag>
                </li>
            </ul>
        </div>
    </div>
</template>
<style scoped></style>