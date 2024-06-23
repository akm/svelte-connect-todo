.PHONY: default
default:

PATH_TO_PROJECT:=../../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

ifeq ($(GOOGLE_CLOUD_PROJECT),)
.PHONY: deploy
deploy:
	@echo "GOOGLE_CLOUD_PROJECT is not set. Please set GOOGLE_CLOUD_PROJECT to deploy to GCP."
else
.PHONY: deploy
deploy: container_push
	gcloud run deploy $(APP_CONTAINER_IMAGE_NAME) \
	--image=$(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) \
	--allow-unauthenticated \
	--port=8080 \
	--memory=128Mi \
	--use-http2 \
	--region=$(GCP_REGION) \
	--project=$(GOOGLE_CLOUD_PROJECT)
endif

.PHONY: update-traffic
update-traffic:
	gcloud run services update-traffic $(APP_CONTAINER_IMAGE_NAME) \
		--to-latest \
		--platform=managed \
		--region=$(GCP_REGION) \
		--project=$(GOOGLE_CLOUD_PROJECT)

.PHONY: inspect_service
inspect_service:
	gcloud run services describe $(APP_CONTAINER_IMAGE_NAME) --region=$(GCP_REGION) --project=$(GOOGLE_CLOUD_PROJECT)

GCP_SERVICE_URL=$(shell gcloud run services describe $(APP_CONTAINER_IMAGE_NAME) --region=$(GCP_REGION) --project=$(GOOGLE_CLOUD_PROJECT) --format=json | jq -r .status.url)
.PHONY: service_url
service_url:
	@echo $(GCP_SERVICE_URL)

GCP_SERVICE_DOMAIN=$(shell echo $(GCP_SERVICE_URL) | sed -e 's/https:\/\///' | sed -e 's/\/.*//')
.PHONY: service_domain
service_domain:
	@echo $(GCP_SERVICE_DOMAIN)

.PHONY: test-connections
test-connections:
	GCP_SERVICE_DOMAIN=$(GCP_SERVICE_DOMAIN) ./test-connections.sh
