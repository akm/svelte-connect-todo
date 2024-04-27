import { browser } from '$app/environment';

export const apisvrOrigin =
	(browser
		? import.meta.env.VITE_APISVR_ORIGIN
		: import.meta.env.VITE_APISVR_ORIGIN_FROM_SERVER || import.meta.env.VITE_APISVR_ORIGIN) ??
	'http://localhost:8080';
