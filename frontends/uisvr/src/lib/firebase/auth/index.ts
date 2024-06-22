import { env } from '$env/dynamic/public';
import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from '../app';

export const auth = getAuth(app);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (env.SK_PUBLIC_FIREBASE_AUTH_EMULATOR_HOST) {
	connectAuthEmulator(auth, 'http://' + env.SK_PUBLIC_FIREBASE_AUTH_EMULATOR_HOST);
}

export type { UserCredential } from 'firebase/auth';

// aliases
export {
	createUserWithEmailAndPassword,
	signInWithEmailAndPassword,
	updateProfile
} from 'firebase/auth';
