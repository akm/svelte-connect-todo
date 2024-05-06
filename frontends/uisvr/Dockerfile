# https://cloud.google.com/run/docs/quickstarts/deploy-continuously?hl=ja#cloudrun_deploy_continuous_code-go
# https://github.com/GoogleCloudPlatform/cloud-run-microservice-template-go/blob/main/Dockerfile
# https://kit.svelte.dev/docs/adapter-node

FROM node:20-bookworm as builder

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /srv/uisvr
COPY package*.json ./
RUN npm install
COPY . ./
RUN npm run build
RUN npm ci --omit dev

FROM node:20-bookworm-slim

WORKDIR /srv/uisvr
COPY --from=builder /srv/uisvr /srv/uisvr

# https://kit.svelte.dev/docs/adapter-node
CMD ["node", "-r", "dotenv/config", "build"]
