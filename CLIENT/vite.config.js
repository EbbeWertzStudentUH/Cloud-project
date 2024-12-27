import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server:{
		proxy: {
			'/proxy/mongo': 'http://localhost:8081',
			'/proxy/mysql': 'http://localhost:8080',
		},
	}
});
