import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	root: './',
	server: {
		proxy: {
			"/api": {
				target: "http://localhost:4001",
				changeOrigin: true,
			},
			"/ws": {
				target: "http://localhost:4001",
				changeOrigin: true,
			},
		},
	},
});
