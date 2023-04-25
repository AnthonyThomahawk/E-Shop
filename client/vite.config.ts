import { defineConfig } from 'vite'
// @ts-ignore
import { sveltekit } from "@sveltejs/kit/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [sveltekit()],
  server: {
    fs: {
      strict: false
    }
  }
})
