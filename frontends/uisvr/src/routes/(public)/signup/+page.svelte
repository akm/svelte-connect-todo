<script lang="ts">
	import { page } from '$app/stores';
	import Icon from '@iconify/svelte';

	import { isFirebaseError } from '$lib/firebase';
	import type { UserCredential } from '$lib/firebase/auth';
	import { auth, createUserWithEmailAndPassword, updateProfile } from '$lib/firebase/auth';

	let email = '';
	let password = '';
	let accountName = '';
	let errorMessage = '';

	const signup = async () => {
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

	const signupOnEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter') {
			signup();
		}
	};
	const clearErrorMessage = () => {
		errorMessage = '';
	};
</script>

<div class="prose space-y-4 m-8">
	<h1 class="mb-4">Sign up</h1>

	{#if errorMessage}
		<div class="flex">
			<Icon icon="exclamation-mark-fill" />
			<p class="text-red-500 ml-2">
				{errorMessage}
			</p>
		</div>
	{/if}

	<label>
		<span>Email</span>
		<input
			class="mt-1 block w-full"
			bind:value={email}
			placeholder="your email address"
			on:keypress={signupOnEnter}
			on:change={clearErrorMessage}
		/>
	</label>

	<label class="block">
		<span>Passeord</span>
		<input
			class="mt-1 block w-full"
			type="password"
			bind:value={password}
			placeholder="your password"
			on:keypress={signupOnEnter}
			on:change={clearErrorMessage}
		/>
	</label>

	<label class="block">
		<span>Account name</span>
		<input
			class="mt-1 block w-full"
			bind:value={accountName}
			placeholder="your account name"
			on:keypress={signupOnEnter}
			on:change={clearErrorMessage}
		/>
		<p class="text-sm">You can change your account name after sign up</p>
	</label>

	<div class="flex">
		<button class="btn btn-primary flex-none mt-4 h-12" on:click={signup}>Sign up</button>
		<div>
			<a class="btn btn-neutral mt-4 ml-8 h-12" color="alternative" href="/signin">Sign in</a>
			<div class="ml-8">Sign in if you already have your account.</div>
		</div>
	</div>
</div>
