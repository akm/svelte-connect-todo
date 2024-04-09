<script lang="ts">
	import type { Task } from '$lib/models/task';
	import type { ServerLoadEvent } from '@sveltejs/kit';
	import { TaskService } from '../gen/task/v1/task_connect';
	import { createPromiseClient } from '@connectrpc/connect';
	import { createConnectTransport } from '@connectrpc/connect-web';
	import { Status } from '../gen/task/v1/task_pb';

	export let data: { tasks: Task[] };

	const transport = createConnectTransport({ baseUrl: 'http://localhost:8080' });
	const client = createPromiseClient(TaskService, transport);
</script>

<div class="centered">
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
							await client.update({ id: task.id, status: done ? Status.DONE : Status.TODO });
						}}
					/>
					<span>{task.name}</span>
					<button
						aria-label="Mark as complete"
						on:click={async (e) => {
							await client.delete({ id: task.id });
							data.tasks = data.tasks.filter((t) => t !== task);
						}}
					/>
				</label>
			</li>
		{/each}
	</ul>
</div>

<style>
	.centered {
		max-width: 20em;
		margin: 0 auto;
	}

	label {
		display: flex;
		width: 100%;
	}

	input[type='text'] {
		flex: 1;
		color: black;
	}

	span {
		flex: 1;
	}

	button {
		border: none;
		background: url(./remove.svg) no-repeat 50% 50%;
		background-size: 1rem 1rem;
		cursor: pointer;
		height: 100%;
		aspect-ratio: 1;
		opacity: 0.5;
		transition: opacity 0.2s;
	}

	button:hover {
		opacity: 1;
	}
</style>
