<script lang="ts">
	export let onEnter: { (name: string): Promise<void> };
</script>

<label>
	add a todo:
	<input
		type="text"
		autocomplete="off"
		style="background: red;"
		on:keydown={async (e) => {
			if (e.key === 'Enter') {
				onEnter;
				const input = e.currentTarget;
				const name = input.value;
				await onEnter(name);

				input.value = '';
			}
		}}
	/>
</label>
