import { initializeApp, getApp, getApps } from 'firebase-admin/app';
import { getAuth } from 'firebase-admin/auth';

import { firebaseConfig } from '$lib/firebase/firebaseconfig';

if (!process.env.GOOGLE_CLOUD_PROJECT && import.meta.env.VITE_GOOGLE_CLOUD_PROJECT) {
	process.env.GOOGLE_CLOUD_PROJECT = import.meta.env.VITE_GOOGLE_CLOUD_PROJECT;
}
// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#admin_sdks
const firebaseAuthEmulatorHost =
	import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST_FROM_SERVER ||
	import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST;
console.log(
	'src/lib/server/firebase-admin/index.ts firebaseAuthEmulatorHost',
	firebaseAuthEmulatorHost
);
if (!process.env.FIREBASE_AUTH_EMULATOR_HOST && firebaseAuthEmulatorHost) {
	process.env.FIREBASE_AUTH_EMULATOR_HOST = firebaseAuthEmulatorHost;
}
console.log(
	'src/lib/server/firebase-admin/index.ts process.env.FIREBASE_AUTH_EMULATOR_HOST',
	process.env.FIREBASE_AUTH_EMULATOR_HOST
);

// FirebaseAppError: The default Firebase app already exists.
// This means you called initializeApp() more than once without
// providing an app name as the second argument.
// というエラーが npm run dev での実行時に発生していたので単純に
// その場合は initializeApp() を呼ぶのではなく、apps.length で判定し、
// 存在していた場合には app() を呼ぶように修正
// See https://vitejs.dev/guide/env-and-mode
const adminApp = (() => {
	const config = { projectId: firebaseConfig.projectId };
	if (import.meta.env.DEV) {
		return getApps().length === 0 ? initializeApp(config) : getApp();
	} else {
		return initializeApp(config);
	}
})();

export const auth = getAuth(adminApp);

export type { DecodedIdToken } from 'firebase-admin/auth';
