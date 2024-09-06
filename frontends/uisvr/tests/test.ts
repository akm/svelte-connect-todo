import { test, expect } from '@playwright/test';

test('test', async ({ page }) => {
	await page.waitForTimeout(5_000);
	await page.goto('http://localhost:8001/signin');
	await expect(page.getByRole('heading', { name: 'TODOs' })).toBeVisible();
	await expect(page.getByRole('heading', { name: 'Sign in' })).toBeVisible();
	await expect(page.getByPlaceholder('your email address')).toBeVisible();
	await expect(page.getByPlaceholder('your password')).toBeVisible();
	await expect(page.getByRole('button', { name: 'Sign in' })).toBeVisible();
	await expect(page.getByRole('link', { name: 'Sign up' })).toBeVisible();
	await page.getByRole('link', { name: 'Sign up' }).click();
	await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
	await expect(page.getByPlaceholder('your email address')).toBeVisible();
	await expect(page.getByPlaceholder('your password')).toBeVisible();
	await expect(page.getByPlaceholder('your account name')).toBeVisible();
	await expect(page.getByRole('button', { name: 'Sign up' })).toBeVisible();
	await expect(page.getByRole('main').getByRole('link', { name: 'Sign in' })).toBeVisible();
	await page.getByRole('button', { name: 'Sign up' }).click();
	await expect(page.getByText('[auth/invalid-email] Firebase')).toBeVisible();
	await page.getByPlaceholder('your email address').click();
	await page.getByPlaceholder('your email address').fill('foo@example.com');
	await page.getByPlaceholder('your password').click();
	await page.getByPlaceholder('your password').fill('Passw0rd!');
	await page.getByPlaceholder('your password').press('Tab');
	await page.getByPlaceholder('your account name').fill('FOO');
	await page.getByRole('button', { name: 'Sign up' }).click();
	await expect(page.getByRole('heading', { name: 'TODOs' })).toBeVisible();
	await expect(page.getByText('add a todo:')).toBeVisible();
	await expect(page.getByText('List items')).toBeVisible();
	await expect(page.getByText('Buy items')).toBeVisible();
	await expect(page.getByLabel('List items')).toBeChecked();
	await expect(page.getByLabel('Buy items')).not.toBeChecked();
	await page.getByLabel('add a todo:').click();
	await page.getByLabel('add a todo:').fill('Sell Items');
	await page.getByLabel('add a todo:').press('Enter');
	await expect(page.getByLabel('Sell Items')).not.toBeChecked();
	await page.getByLabel('Buy items').check();
	await expect(page.getByLabel('Buy items')).toBeChecked();
	await page.getByLabel('List items').uncheck();
	await expect(page.getByLabel('List items')).not.toBeChecked();
	await page
		.locator('label')
		.filter({ hasText: 'List items' })
		.getByLabel('Mark as complete')
		.click();
	await expect(page.getByText('List items')).not.toBeVisible();
	await page.getByRole('button', { name: 'Sign out' }).click();
});
