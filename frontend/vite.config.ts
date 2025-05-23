import path from "path"
import react from "@vitejs/plugin-react"
import { defineConfig, loadEnv } from "vite"

export default defineConfig(({ command, mode }) => {
  console.log('latest base')
  const env = loadEnv(mode, process.cwd())
  return {
    base: "",
    plugins: [react()],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    server: mode === 'development' ? {
      proxy: {
        '/patlytics': {
          target: env.VITE_DEV_API_URL,
          secure: false,
          changeOrigin: true,
        }
      }
    } : {}
  }
})
