import type { FirebaseError } from 'firebase/app';

export const isFirebaseError = (error: unknown): error is FirebaseError => {
	return error instanceof Error && 'code' in error && 'message' in error;
};

export type { FirebaseError } from 'firebase/app';
