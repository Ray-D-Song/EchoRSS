import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import generouted from '@generouted/react-router/plugin'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), generouted()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
    },
  },
});
