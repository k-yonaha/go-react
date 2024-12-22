import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',  // 任意のIPからアクセス可能にする
    port: 5173         // Viteの開発サーバーのポート
  }
})
