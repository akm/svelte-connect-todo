<script lang="ts">
	import Icon from '@iconify/svelte';

	let email = '';
	let password = '';

	export let errorMessage = '';
	export let onSignin: { (email: string, password: string): Promise<void> };

	const signin = () => {
		return onSignin(email, password);
	};

	const signinOnEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter') {
			signin();
		}
	};
	const clearErrorMessage = () => {
		errorMessage = '';
	};
</script>

<div class="prose space-y-4 m-8">
	<h1 class="mb-4">Sign in</h1>

	{#if errorMessage}
		<div class="flex">
			<Icon icon="exclamation-mark-fill" />
			<p class="text-red-500 ml-2">
				{errorMessage}
			</p>
		</div>
	{/if}

	<label class="block">
		<span>Email</span>
		<input
			class="mt-1 block w-full"
			bind:value={email}
			placeholder="your email address"
			on:keypress={signinOnEnter}
			on:change={clearErrorMessage}
		/>
	</label>

	<label class="block">
		<span>Password</span>
		<input
			class="mt-1 block w-full"
			type="password"
			bind:value={password}
			placeholder="your password"
			on:keypress={signinOnEnter}
			on:change={clearErrorMessage}
		/>
	</label>
	<div class="flex">
		<button class="btn btn-primary flex-none mt-4 h-12" on:click={signin}>Sign in</button>
		<div>
			<a class="btn btn-neutral mt-4 ml-8 h-12" color="alternative" href="/signup">Sign up</a>
			<div class="ml-8">Sign up if you don't have an account.</div>
		</div>
	</div>
</div>
