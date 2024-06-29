# Local development stage - stages/local/dev

## How to run

```
make -C stages/local/dev up
```

and Open http://localhost:8000/ in browser.

## How to shutdown

```
make -C stages/local/dev down
```

## How to develop

Set `rproxy`, `uisvr` or `apisvr` to environment variable `DEV_TARGET`.
You can launch the server manually on host instead of container.

When you want to launch containers for uisvr, Run the following;

```
DEV_TARGET=uisvr make -C stages/local/dev up
```

When you want to debug uisvr, just do this

```
make -C frontends/uisvr dev
```

( This means `cd frontends/uisvr && make dev` )
