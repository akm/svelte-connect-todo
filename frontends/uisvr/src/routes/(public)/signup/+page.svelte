<script lang="ts">
	import { page } from '$app/stores';

	import SignupForm from '$lib/components/forms/SignupForm.svelte';

	import { isFirebaseError } from '$lib/firebase';
	import type { UserCredential } from '$lib/firebase/auth';
	import { auth, createUserWithEmailAndPassword, updateProfile } from '$lib/firebase/auth';

	let errorMessage = '';

	const signup = async (email: string, password: string, accountName: string) => {
		let credential: UserCredential;
		try {
			credential = await createUserWithEmailAndPassword(auth, email, password);
		} catch (err) {
			if (isFirebaseError(err)) {
				errorMessage = `[${err.code}] ${err.message}`;
				return;
			} else {
				throw err;
			}
		}
		if (accountName != '') {
			try {
				await updateProfile(credential.user, { displayName: accountName });
			} catch (err) {
				console.error(err);
				// 例外をthrowしてしまうと、再度SignUpするためにはユーザーを削除するか、ユーザーが
				// 登録済みかどうかの判断が必要になります。基本的に displayName は後でユーザー
				// が変更可能なものなので、ここで設定に失敗した場合には 「未設定」 と表示すること
				// にして、ここではエラーにしません。
				// throw err;
			}
		}

		const idToken = await credential.user.getIdToken();
		await fetch('/session', {
			method: 'POST',
			body: JSON.stringify({ id_token: idToken }),
			headers: { 'Content-Type': 'application/json' }
		});

		window.location.href = $page.url.origin + '/';
	};
</script>

<SignupForm bind:errorMessage onSignup={signup} />
