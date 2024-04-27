import { browser } from '$app/environment';
import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from '../app';

export const auth = getAuth(app);

console.log('src/lib/firebase/auth/index.ts browser', browser);

const firebaseAuthEmulatorHost = browser
	? import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST
	: import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST_FROM_SERVER ||
		import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST;

console.log('src/lib/firebase/auth/index.ts firebaseAuthEmulatorHost', firebaseAuthEmulatorHost);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (firebaseAuthEmulatorHost) {
	connectAuthEmulator(auth, 'http://' + firebaseAuthEmulatorHost);
}

export type { UserCredential } from 'firebase/auth';

// aliases
export {
	createUserWithEmailAndPassword,
	signInWithEmailAndPassword,
	updateProfile
} from 'firebase/auth';
