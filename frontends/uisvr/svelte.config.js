import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// See https://kit.svelte.dev/docs/adapters
		//     https://kit.svelte.dev/docs/adapter-node
		adapter: adapter(),
		// https://kit.svelte.dev/docs/configuration#env
		env: {
			publicPrefix: 'SK_PUBLIC_',
			privatePrefix: 'SK_PRIVATE_'
		}
	}
};

export default config;
