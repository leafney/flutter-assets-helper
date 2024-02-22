<script setup lang="ts">
import { ref } from "vue";

const colorValue = ref('')


const copyColor = async () => {
    try {
        let copyText = await navigator.clipboard.readText();
        if (copyText) {
            colorValue.value = copyText;
        } else {
            console.log('剪切板为空');
        }
    } catch (error) {
        console.error(error);
    }
}

</script>
<template>
    <div>
        <button @click="copyColor">粘贴色值</button>
        <h3>{{ colorValue }}</h3>
    </div>
</template>
<style scoped></style>