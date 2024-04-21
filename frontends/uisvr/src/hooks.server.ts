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
		// {
		// 	email: 'foo2@example.com',
		// 	email_verified: false,
		// 	auth_time: 1713688711,
		// 	user_id: 'z627DaZCAl6ESgxlf76JXFYTraN5',
		// 	firebase: { identities: { email: [Array] }, sign_in_provider: 'password' },
		// 	iat: 1713688711,
		// 	exp: 1714120711,
		// 	aud: 'govelte-app1-gcp-project1',
		// 	iss: 'https://session.firebase.google.com/govelte-app1-gcp-project1',
		// 	sub: 'z627DaZCAl6ESgxlf76JXFYTraN5',
		// 	uid: 'z627DaZCAl6ESgxlf76JXFYTraN5'
		// }
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
