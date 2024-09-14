<script lang="ts">
	import { page } from '$app/stores';

	import SigninForm from '$lib/components/forms/SigninForm.svelte';

	import { isFirebaseError } from '$lib/firebase';
	import type { UserCredential } from '$lib/firebase/auth';
	import { auth, signInWithEmailAndPassword } from '$lib/firebase/auth';

	let errorMessage = '';

	const signin = async (email: string, password: string) => {
		let userCredential: UserCredential;
		try {
			userCredential = await signInWithEmailAndPassword(auth, email, password);
			console.log('userCredential', userCredential);
		} catch (err) {
			if (isFirebaseError(err)) {
				errorMessage = `[${err.code}] ${err.message}`;
				return;
			}
			throw err;
		}
		const idToken = await userCredential.user.getIdToken();
		await fetch('/session', {
			method: 'POST',
			body: JSON.stringify({ id_token: idToken }),
			headers: { 'Content-Type': 'application/json' }
		});
		window.location.href = $page.url.origin + '/';
	};
</script>

<SigninForm bind:errorMessage onSignin={signin} />
