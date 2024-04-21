import type { Handle } from '@sveltejs/kit';
// import { FirebaseError } from 'firebase-admin';

import { isFirebaseError } from '$lib/firebase';
import { auth } from '$lib/server/firebase-admin';
import type { DecodedIdToken } from '$lib/server/firebase-admin';

export const handle: Handle = async ({ event, resolve }) => {
	// get cookies from browser
	const sessionValue = event.cookies.get('session');

	if (!sessionValue) {
		// if there is no session load page as normal
		return await resolve(event);
	}

	let decoded: DecodedIdToken;
	try {
		decoded = await auth.verifySessionCookie(sessionValue, true);
		// Contents of decoded
		console.log('decoded', decoded);
		// firebase.identities.email は [ 'foo1@example.com' ]
	} catch (err) {
		if (isFirebaseError(err)) {
			switch (err.code) {
				case 'auth/session-cookie-revoked':
				case 'auth/user-not-found':
					event.locals.user = undefined;
					return await resolve(event);
			}
		}
		console.error('verifySessionCookie error', err);
		throw err;
	}
	const user = await auth.getUser(decoded.uid);
	event.locals.user = {
		id: user.uid,
		name: user.displayName,
		email: user.email
	};

	// load page as normal
	return await resolve(event);
};
