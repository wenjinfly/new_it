import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      // 带选项写法：http://localhost:5173/api/bar -> http://127.0.0.1:12345/bar
      '/api': {
        target: 'http://127.0.0.1:12345',// 后端接口
        changeOrigin: true, // 是否跨域
        rewrite: (path) => path.replace(/^\/api/, ''),
      }
    }
  }
})
