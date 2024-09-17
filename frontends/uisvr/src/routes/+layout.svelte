<script>
	import 'tailwindcss/tailwind.css';
	import Header from '$lib/components/compounds/Header.svelte';

	import { page } from '$app/stores';
	const user = $page.data.user;

	const signout = async () => {
		await fetch('/session', { method: 'DELETE' });
		window.location.href = $page.url.origin + '/signin';
	};
</script>

<Header {user} onClickSignOut={signout} />

<main class="container mx-auto mt-8">
	<slot />
</main>
