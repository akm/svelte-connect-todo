<script>
	import 'tailwindcss/tailwind.css';
	import Header from '$lib/components/compounds/Header.svelte';
</script>

<Header />

<main class="container mx-auto mt-8">
	<slot />
</main>
