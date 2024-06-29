# production stage (or staging stage)

## Setup overview

1. Deploy apisvr (without environment variables for CORS)
2. Deploy uisvr
3. Deploy apisvr again (with environment variables for CORS)

## How to setup on CloudRun

```
make -C stages/production/apisvr deploy
```

```
make -C stages/production/apisvr test-connections
```

```
make -C stages/production/uisvr deploy
```

```
make -C stages/production/apisvr deploy
```

```
make -C stages/production/uisvr service_url
```

The URL of uisvr will be shown. Open the URL in browser.
