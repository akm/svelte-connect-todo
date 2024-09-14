<script lang="ts">
	export let user: { name: string };
	export let onClickSignOut: { (): Promise<void> };
</script>

<header class="bg-gray-800 text-white p-4">
	<div class="container mx-auto flex justify-between items-center">
		<h1 class="text-lg font-bold">TODOs</h1>
		<div class="text-sm">
			{#if user}
				<span class="font-medium">{user.name}</span>
				<button class="btn ml-4" on:click={onClickSignOut}>Sign out</button>
			{:else}
				<a href="/signin" class="hover:underline">Sign in</a>
			{/if}
		</div>
	</div>
</header>
