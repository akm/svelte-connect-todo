.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

include $(PATH_TO_SHAPEAPPMK)/git/current.mk

DOCKER_IMAGE_REPOSITORY=$(shell $(MAKE) -C $(PATH_TO_APISVR) --no-print-directory docker-image-name)
DOCKER_IMAGE_NAME_COMMIT_HASH=$(DOCKER_IMAGE_REPOSITORY):$(GIT_CURRENT_COMMIT_HASH_SHORT)
DOCKER_IMAGE_NAME?=$(DOCKER_IMAGE_REPOSITORY):$(APP_STAGE_TYPE)
DOCKER_IMAGE_BUILD_OPTS=\
	--platform linux/amd64 \
	--build-arg APP_BIN_PATH=$(shell $(MAKE) golang-binary-path-for-stage-production --no-print-directory) \
	-f Dockerfile
include $(PATH_TO_SHAPEAPPMK)/docker/image.mk

.PHONY: build
build:
	cd $(PATH_TO_APISVR) && \
	$(MAKE) golang-binary-build-for-stage-production && \
	$(MAKE) docker-image-build

.PHONY: push
push: build
	docker tag $(DOCKER_IMAGE_NAME) $(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) && \
	docker push $(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_NAME_COMMIT_HASH)

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
	--image=$(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) \
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

GOOGLE_CLOUD_RUN_SERVICE_NAME=$(DOCKER_IMAGE_REPOSITORY)

ifeq ($(GOOGLE_CLOUD_PROJECT),)
.PHONY: deploy
deploy:
	@echo "GOOGLE_CLOUD_PROJECT is not set. Please set GOOGLE_CLOUD_PROJECT to deploy to GCP."
else
.PHONY: deploy
deploy: push
	$(GCLOUD) run deploy $(GOOGLE_CLOUD_RUN_SERVICE_NAME) $(shell $(MAKE) deploy_options --no-print-directory)
endif

.PHONY: update-traffic
update-traffic:
	$(GCLOUD) run services update-traffic $(GOOGLE_CLOUD_RUN_SERVICE_NAME) \
		--to-latest \
		--platform=managed \
		--region=$(GCP_REGION)

.PHONY: inspect_service
inspect_service:
	$(GCLOUD) run services describe $(GOOGLE_CLOUD_RUN_SERVICE_NAME) --region=$(GCP_REGION)

GCP_SERVICE_URL=$(shell $(GCLOUD) run services describe $(GOOGLE_CLOUD_RUN_SERVICE_NAME) --region=$(GCP_REGION) --format='value(status.url)' )
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
