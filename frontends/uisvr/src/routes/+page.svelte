<script lang="ts">
	import type { Task } from '$lib/models/task';
	import { TaskService } from '../gen/task/v1/task_connect';
	import { createPromiseClient } from '@connectrpc/connect';
	import { createConnectTransport } from '@connectrpc/connect-web';
	import { TaskStatus } from '../gen/task/v1/task_pb';
	import { apisvrOrigin } from '$lib/apisvr';

	import TaskForm from '$lib/components/forms/TaskForm.svelte';
	import TaskList from '$lib/components/collections/TaskList.svelte';

	export let data: { tasks: Task[] };

	const transport = createConnectTransport({
		baseUrl: apisvrOrigin,
		credentials: 'include'
	});
	const client = createPromiseClient(TaskService, transport);
</script>

<div class="prose m-8 lg:prose-lg">
	<TaskForm
		onEnter={async (name) => {
			const response = await client.create({ name, status: TaskStatus.TODO });
			const { id } = response;
			data.tasks = [...data.tasks, { id, name, done: false }];
		}}
	/>

	<TaskList
		bind:tasks={data.tasks}
		onClickCheckbox={async (task, checked) => {
			await client.update({
				id: task.id,
				name: task.name,
				status: checked ? TaskStatus.DONE : TaskStatus.TODO
			});
		}}
		onClickDelete={async (task) => {
			await client.delete({ id: task.id });
			data.tasks = data.tasks.filter((t) => t !== task);
		}}
	/>
</div>
