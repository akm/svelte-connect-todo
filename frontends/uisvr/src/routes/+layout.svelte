<script>
  import "tailwindcss/tailwind.css";
	import { page } from '$app/stores';
  const user = $page.data.user;
</script>

<header class="bg-gray-800 text-white p-4">
  <div class="container mx-auto flex justify-between items-center">
      <h1 class="text-lg font-bold">TODOs</h1>
      <div class="text-sm">
        {#if user}
          <span class="font-medium">{user.name}</span>
        {:else}
          <a href="/signin" class="hover:underline">Sign in</a>
        {/if}
      </div>
  </div>
</header>

<main class="container mx-auto mt-8">

  <slot />

</main>
