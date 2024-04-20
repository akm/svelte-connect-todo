import type { RequestEvent } from '@sveltejs/kit';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function POST(event: RequestEvent) {
	const { request, cookies } = event;
	const { id_token } = await request.json();
	if (!id_token) throw new Error('id_token not found');
	cookies.set('session', id_token, { path: '/', httpOnly: true, sameSite: 'lax' });
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
