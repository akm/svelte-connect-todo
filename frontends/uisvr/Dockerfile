# https://cloud.google.com/run/docs/quickstarts/deploy-continuously?hl=ja#cloudrun_deploy_continuous_code-go
# https://github.com/GoogleCloudPlatform/cloud-run-microservice-template-go/blob/main/Dockerfile
# https://kit.svelte.dev/docs/adapter-node

FROM node:20-bookworm as builder

ARG VITE_GOOGLE_CLOUD_PROJECT
ARG VITE_FIREBASE_AUTH_EMULATOR_HOST

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /srv/uisvr
COPY package*.json ./
RUN npm install
COPY . ./

ENV VITE_GOOGLE_CLOUD_PROJECT=${VITE_GOOGLE_CLOUD_PROJECT}
ENV VITE_FIREBASE_AUTH_EMULATOR_HOST=${VITE_FIREBASE_AUTH_EMULATOR_HOST}

RUN npm ci --omit dev

CMD ["node", "build"]
