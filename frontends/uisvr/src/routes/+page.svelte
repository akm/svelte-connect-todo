<script lang="ts">
	import type { Task } from '$lib/models/task';
	import { TaskService } from '../gen/task/v1/task_connect';
	import { createPromiseClient } from '@connectrpc/connect';
	import { createConnectTransport } from '@connectrpc/connect-web';
	import { Status } from '../gen/task/v1/task_pb';
	import Icon from '@iconify/svelte';

	export let data: { tasks: Task[] };

	const transport = createConnectTransport({ baseUrl: 'http://localhost:8080' });
	const client = createPromiseClient(TaskService, transport);
</script>

<div class="prose m-8 lg:prose-lg">
	<h1>todos</h1>

	<label>
		add a todo:
		<input
			type="text"
			autocomplete="off"
			on:keydown={async (e) => {
				if (e.key === 'Enter') {
					const input = e.currentTarget;
					const name = input.value;

					const response = await client.create({ name, status: Status.TODO });
					const { id } = response;
					data.tasks = [...data.tasks, { id, name, done: false }];

					input.value = '';
				}
			}}
		/>
	</label>

	<ul class="todos">
		{#each data.tasks as task (task.id)}
			<li>
				<label>
					<input
						type="checkbox"
						checked={task.done}
						on:change={async (e) => {
							const done = e.currentTarget.checked;
							await client.update({ id: task.id, name: task.name, status: done ? Status.DONE : Status.TODO });
						}}
					/>
					<span>{task.name}</span>
					<button
						aria-label="Mark as complete"
						on:click={async (e) => {
							await client.delete({ id: task.id });
							data.tasks = data.tasks.filter((t) => t !== task);
						}}
					><Icon icon="ph:trash-light" /></button>
				</label>
			</li>
		{/each}
	</ul>
</div>
