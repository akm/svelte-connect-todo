import { env } from '$env/dynamic/public';

// For Firebase JS SDK v7.20.0 and later, measurementId is optional
export const fbConfig = {
	apiKey: env.SK_PUBLIC_FB_CONFIG_API_KEY || 'dummy-firebase-api-key1',
	authDomain: env.SK_PUBLIC_FB_CONFIG_AUTH_DOMAIN || 'svelte-connect-todo.firebaseapp.com',
	projectId: env.SK_PUBLIC_FB_CONFIG_PROJECT_ID || 'svelte-connect-todo-gcp-project1',
	storageBucket: env.SK_PUBLIC_FB_CONFIG_STORAGE_BUCKET || 'svelte-connect-todo.appspot.com',
	messagingSenderId: env.SK_PUBLIC_FB_CONFIG_MESSAGING_SENDER_ID || '012345678901',
	appId: env.SK_PUBLIC_FB_CONFIG_APP_ID || '1:012345678901:web:0123456789abcdefghijkl',
	measurementId: env.SK_PUBLIC_FB_CONFIG_MEASUREMENT_ID || 'G-ZZZZZZZZZZ'
};
