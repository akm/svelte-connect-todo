import { defineConfig } from '@playwright/test';

// https://playwright.dev/docs/browsers
// https://playwright.dev/docs/api/class-testconfig#test-config-projects

export default defineConfig({
	// timeout: 5 * 60_000,
	use: {
		headless: process.env.HEADED != 'true',
		launchOptions: {
			slowMo: Number(process.env.SLOMO || 0) // テスト実行時のスローモーション。デモなら 2000 くらいがよい
		},
		baseURL: process.env.APP_BASE_URL ?? 'http://localhost:8001',
		trace: 'retain-on-failure',

		// https://playwright.dev/docs/api/class-testoptions#test-options-video
		video: process.env.VIDEO === 'on' ? 'on' : 'off'
	},

	// https://playwright.dev/docs/api/class-testconfig#test-config-web-server
	webServer: {
		command: 'make -C ../../stages/localtest run',
		port: 8001,
		// Github Actions で実行する場合は stages/localtest の セットアップも行うので 5 分くらいかかることもありえる
		timeout: 300_000
	},

	testDir: 'tests',
	testMatch: /(.+\.)?(test|spec)\.[jt]s/,
	workers: 1,

	// https://playwright.dev/docs/test-reporters#github-actions-annotations
	reporter: process.env.CI ? 'github' : 'list'
});
