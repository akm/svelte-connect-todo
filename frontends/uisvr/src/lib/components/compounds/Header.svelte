<script lang="ts">
	import { page } from '$app/stores';
	const user = $page.data.user;

	const signOut = async () => {
		await fetch('/session', { method: 'DELETE' });
		window.location.href = $page.url.origin + '/signin';
	};
</script>

<header class="bg-gray-800 text-white p-4">
	<div class="container mx-auto flex justify-between items-center">
		<h1 class="text-lg font-bold">TODOs</h1>
		<div class="text-sm">
			{#if user}
				<span class="font-medium">{user.name}</span>
				<button class="btn ml-4" on:click={signOut}>Sign out</button>
			{:else}
				<a href="/signin" class="hover:underline">Sign in</a>
			{/if}
		</div>
	</div>
</header>
