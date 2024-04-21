import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from '..';

export const auth = getAuth(app);

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
