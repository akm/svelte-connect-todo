# https://cloud.google.com/run/docs/quickstarts/deploy-continuously?hl=ja#cloudrun_deploy_continuous_code-go
# https://github.com/GoogleCloudPlatform/cloud-run-microservice-template-go/blob/main/Dockerfile

FROM node:20-bookworm as builder
WORKDIR /srv/uisvr
COPY package*.json ./
RUN npm install
COPY . ./
RUN npm ci --omit dev


FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /srv/uisvr /srv/uisvr
WORKDIR /srv/uisvr

CMD ["node buill"]
