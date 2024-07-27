# https://cloud.google.com/run/docs/quickstarts/deploy-continuously?hl=ja#cloudrun_deploy_continuous_code-go
# https://github.com/GoogleCloudPlatform/cloud-run-microservice-template-go/blob/main/Dockerfile

FROM debian:bookworm-slim

ARG APP_BIN_PATH=./bin/server-local

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY ${APP_BIN_PATH} /srv/server

CMD ["/srv/server"]
