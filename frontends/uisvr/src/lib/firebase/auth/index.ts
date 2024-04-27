import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from '../app';

export const auth = getAuth(app);

console.log(
	'src/lib/firebase/auth/index.ts import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST',
	import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST
);
console.log(
	'src/lib/firebase/auth/index.ts process.env.VITE_FIREBASE_AUTH_EMULATOR_HOST',
	process.env.VITE_FIREBASE_AUTH_EMULATOR_HOST
);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST) {
	connectAuthEmulator(auth, 'http://' + import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST);
}

export type { UserCredential } from 'firebase/auth';

// aliases
export {
	createUserWithEmailAndPassword,
	signInWithEmailAndPassword,
	updateProfile
} from 'firebase/auth';
