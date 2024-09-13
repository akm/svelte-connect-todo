# Local development stage - stages/local

## How to run

```
make -C stages/local up
```

and Open http://localhost:8000/ in browser.

## How to shutdown

```
make -C stages/local down
```

## How to develop

Set `rproxy`, `uisvr` or `apisvr` to environment variable `DEV_TARGET`.
You can launch the server manually on host instead of container.

When you want to launch containers for uisvr, Run the following;

```
DEV_TARGET=uisvr make -C stages/local up
```

When you want to debug uisvr, just do this

```
make -C frontends/uisvr dev
```

( This means `cd frontends/uisvr && make dev` )

## ports

| DEV_TARGET | fb-emu auth | fb-emu suite | apisvr | uisvr          | rproxy  | apisvr origin from uisvr | apisvr CORS                                 |
| ---------- | ----------- | ------------ | ------ | -------------- | ------- | ------------------------ | ------------------------------------------- |
| all        | 9099        | 4000         | 8080   | 4173           | 8000    | (same origin)            | http://localhost:4173,http://localhost:5173 |
| rproxy     | 9099        | 4000         | 8080   | 4173           | (8000 ) | (same origin)            | http://localhost:4173,http://localhost:5173 |
| uisvr      | 9099        | 4000         | 8080   | (5173 or 4173) | -       | http://localhost:8080    | http://localhost:4173,http://localhost:5173 |
| apisvr     | 9099        | 4000         | (8080) | -              | -       | http://localhost:8080    | http://localhost:4173,http://localhost:5173 |

`(port)` is opened by debugging environment.

- fb-emu auth: Firebase Emulators Authentication
- fb-emu suite: Firebase Emulators Suite
