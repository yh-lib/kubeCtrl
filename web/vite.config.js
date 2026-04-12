import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  // 指定env配置文件目录
  envDir: './src/env',
  // 禁用依赖预构建输出的 source map，避免空的 source map 引起浏览器报错
  optimizeDeps: {
    esbuildOptions: {
      sourcemap: false
    }
  },
  // 构建时也关闭 sourcemap
  build: {
    sourcemap: false
  }
})
