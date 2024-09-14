import type { Meta, StoryObj } from '@storybook/svelte';

import 'tailwindcss/tailwind.css';
import TaskList from '$lib/components/collections/TaskList.svelte';

const meta = {
	title: 'collections/TaskList',
	component: TaskList,
	// This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
	tags: ['autodocs'],
	parameters: {
		// More on how to position stories at: https://storybook.js.org/docs/configure/story-layout
		layout: 'fullscreen'
	}
} satisfies Meta<TaskList>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Basic: Story = {
	args: {
		tasks: [
			{
				id: 1,
				name: 'Task 1',
				done: true
			},
			{
				id: 2,
				name: 'Task 2',
				done: false
			},
			{
				id: 3,
				name: 'Long name task 3 qwertyuiopasdfghjklzxcvbnm',
				done: false
			}
		]
	}
};

export const Empty: Story = {
	args: {
		tasks: []
	}
};
