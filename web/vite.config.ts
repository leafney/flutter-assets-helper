import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// 导入路径处理模块
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@cs": path.resolve(__dirname, "./src/components"),
      "@vs": path.resolve(__dirname, "./src/views"),
    },
  }
})
