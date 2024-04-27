import { browser } from '$app/environment';

console.log('src/lib/apisvr/index.ts browser', browser);
console.log(
	'src/lib/apisvr/index.ts import.meta.env.VITE_APISVR_ORIGIN',
	import.meta.env.VITE_APISVR_ORIGIN
);
console.log(
	'src/lib/apisvr/index.ts import.meta.env.VITE_APISVR_ORIGIN_FROM_SERVER',
	import.meta.env.VITE_APISVR_ORIGIN_FROM_SERVER
);

export const apisvrOrigin =
	(browser
		? import.meta.env.VITE_APISVR_ORIGIN
		: import.meta.env.VITE_APISVR_ORIGIN_FROM_SERVER || import.meta.env.VITE_APISVR_ORIGIN) ??
	'http://localhost:8080';

console.log('src/lib/apisvr/index.ts apisvrOrigin', apisvrOrigin);
