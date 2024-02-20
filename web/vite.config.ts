import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// 导入路径处理模块
import path from "path";
// naive ui 自动导入
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@cs": path.resolve(__dirname, "./src/components"),
      "@vs": path.resolve(__dirname, "./src/views"),
    },
  }
})
