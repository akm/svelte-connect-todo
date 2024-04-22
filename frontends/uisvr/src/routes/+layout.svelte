
<script>
  import "tailwindcss/tailwind.css";
</script>

<header class="bg-gray-800 text-white p-4">
  <div class="container mx-auto flex justify-between items-center">
      <h1 class="text-lg font-bold">TODOs</h1>
      <div class="text-sm">
          User: <span class="font-medium">User name</span>
      </div>
  </div>
</header>

<main class="container mx-auto mt-8">

  <slot />

</main>
