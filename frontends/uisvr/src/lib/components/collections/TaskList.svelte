<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { Task } from '$lib/models/task';

	export let tasks: Task[] = [];

	export let onClickCheckbox: { (task: Task, checked: boolean): Promise<void> };
	export let onClickDelete: { (task: Task): Promise<void> };
</script>

<ul class="todos">
	{#each tasks as task (task.id)}
		<li>
			<label>
				<input
					type="checkbox"
					checked={task.done}
					on:change={async (e) => {
						await onClickCheckbox(task, e.currentTarget.checked);
					}}
				/>
				<span>{task.name}</span>
				<button
					aria-label="Mark as complete"
					on:click={async () => {
						await onClickDelete(task);
					}}
				>
					<Icon icon="ph:trash-light" />
				</button>
			</label>
		</li>
	{/each}
</ul>
