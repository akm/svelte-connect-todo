import type { RequestEvent } from '@sveltejs/kit';
import { auth } from '$lib/server/firebase-admin';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function POST(event: RequestEvent) {
	// Skip checking CSRF token because sveltekit has already done it.
	// https://kit.svelte.jp/docs/configuration#csrf
	const { request, cookies } = event;
	const { id_token } = await request.json();
	if (!id_token) throw new Error('id_token not found');
	// https://firebase.google.com/docs/auth/admin/manage-cookies#create_session_cookie
	const expiresIn = 60 * 60 * 24 * 5 * 1000;
	const sessionCookie = await auth.createSessionCookie(id_token, { expiresIn });
	// Set cookie policy for session cookie.
	cookies.set('session', sessionCookie, {
		path: '/',
		httpOnly: true,
		sameSite: 'lax',
		expires: new Date(Date.now() + expiresIn)
	});
	return new Response(null, { status: 204 });
}

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function DELETE(event: RequestEvent) {
	const { cookies } = event;
	cookies.set('session', '', { path: '/', httpOnly: true, sameSite: 'lax' });
	return new Response(null, { status: 204 });
}
