import { defineConfig } from '@playwright/test';

export default defineConfig({
	webServer: {
		command: 'make -C ../../stages/localtest run',
		port: 8001
	},
	testDir: 'tests',
	testMatch: /(.+\.)?(test|spec)\.[jt]s/
});
