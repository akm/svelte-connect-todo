import type { Meta, StoryObj } from '@storybook/svelte';

import 'tailwindcss/tailwind.css';
import SignupForm from '$lib/components/forms/SignupForm.svelte';

const meta = {
	title: 'forms/SignupForm',
	component: SignupForm,
	// This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
	tags: ['autodocs'],
	parameters: {
		// More on how to position stories at: https://storybook.js.org/docs/configure/story-layout
		layout: 'fullscreen'
	}
} satisfies Meta<SignupForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Basic: Story = {
	args: {}
};
