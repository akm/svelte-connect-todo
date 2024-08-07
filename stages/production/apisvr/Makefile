.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

APP_CONTAINER_IMAGE_NAME=$(shell $(MAKE) -C $(PATH_TO_APISVR) --no-print-directory container_image_name)
APP_CONTAINER_IMAGE_PRODUCTION=$(APP_CONTAINER_IMAGE_NAME):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_image_name
container_image_name:
	@echo $(APP_CONTAINER_IMAGE_NAME)

.PHONY: build
build:
	cd $(PATH_TO_APISVR) && \
	$(MAKE) golang-binary-production && \
	docker build \
		--platform linux/amd64 \
		--build-arg APP_BIN_PATH=$(shell $(MAKE) golang-binary-path-for-stage-production --no-print-directory) \
		-t $(APP_CONTAINER_IMAGE_PRODUCTION) \
		-f Dockerfile \
		.

.PHONY: push
push: build
	docker tag $(APP_CONTAINER_IMAGE_PRODUCTION) $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) && \
	docker push $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION)

UISVR_ORIGIN=$(shell $(MAKE) -C ../uisvr --no-print-directory service_url 2>/dev/null || echo "")
.PHONY: uisvr_origin
uisvr_origin:
	@echo $(UISVR_ORIGIN)

ENV_VAR_YAML=./env-vars.yaml
$(ENV_VAR_YAML):
	$(MAKE) env_vars_yaml_gen

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk

.PHONY: env_vars_yaml_gen
env_vars_yaml_gen: $(TEXT_TEMPLATE_CLI)
	UISVR_ORIGIN=$(UISVR_ORIGIN) $(TEXT_TEMPLATE_CLI) $(ENV_VAR_YAML).tmpl > $(ENV_VAR_YAML)

DEPLOY_OPTIONS_BASE=\
	--image=$(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) \
	--allow-unauthenticated \
	--port=$(APP_PORT_APISVR_dev) \
	--memory=128Mi \
	--use-http2 \
	--region=$(GCP_REGION)
ENV_VARS_CONTENT=$(shell yq '... comments=""' $(ENV_VAR_YAML))

.PHONY: deploy_options
deploy_options:
	@$(MAKE) env_vars_yaml_gen >&2 && \
	echo $(DEPLOY_OPTIONS_BASE) && \
	[ -n "$(ENV_VARS_CONTENT)" ] && echo "--env-vars-file=$(ENV_VAR_YAML)" || echo ""

ifeq ($(GOOGLE_CLOUD_PROJECT),)
.PHONY: deploy
deploy:
	@echo "GOOGLE_CLOUD_PROJECT is not set. Please set GOOGLE_CLOUD_PROJECT to deploy to GCP."
else
.PHONY: deploy
deploy: push
	$(GCLOUD) run deploy $(APP_CONTAINER_IMAGE_NAME) $(shell $(MAKE) deploy_options --no-print-directory)
endif

.PHONY: update-traffic
update-traffic:
	$(GCLOUD) run services update-traffic $(APP_CONTAINER_IMAGE_NAME) \
		--to-latest \
		--platform=managed \
		--region=$(GCP_REGION)

.PHONY: inspect_service
inspect_service:
	$(GCLOUD) run services describe $(APP_CONTAINER_IMAGE_NAME) --region=$(GCP_REGION)

GCP_SERVICE_URL=$(shell $(GCLOUD) run services describe $(APP_CONTAINER_IMAGE_NAME) --region=$(GCP_REGION) --format='value(status.url)' )
.PHONY: service_url
service_url:
	@echo $(GCP_SERVICE_URL)

GCP_SERVICE_DOMAIN=$(shell echo $(GCP_SERVICE_URL) | sed -e 's/https:\/\///' | sed -e 's/\/.*//')
.PHONY: service_domain
service_domain:
	@echo $(GCP_SERVICE_DOMAIN)

.PHONY: test-connections
test-connections:
	SERVER_DOMAIN=$(GCP_SERVICE_DOMAIN) $(MAKE) -C $(PATH_TO_APISVR) test-connections
