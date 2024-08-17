.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

.PHONY: build
build:
	$(MAKE) -C $(PATH_TO_APISVR) golang-binary-build-for-stage-production && \
	$(MAKE) -C $(PATH_TO_APISVR) docker-image-build

DOCKER_IMAGE_REPOSITORY=$(shell $(MAKE) -C $(PATH_TO_APISVR) --no-print-directory docker-image-name)
DOCKER_IMAGE_BUILD_OPTS=\
	--platform linux/amd64 \
	--build-arg APP_BIN_PATH=$(shell $(MAKE) golang-binary-path-for-stage-production --no-print-directory) \
	-f Dockerfile
include $(PATH_TO_SHAPEAPPMK)/default/docker/image.mk

UISVR_ORIGIN=$(shell $(MAKE) -C ../uisvr --no-print-directory service_url 2>/dev/null || echo "")
.PHONY: uisvr_origin
uisvr_origin:
	@echo $(UISVR_ORIGIN)

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk
ENV_VAR_YAML=./env-vars.yaml
$(ENV_VAR_YAML): $(TEXT_TEMPLATE_CLI)
	UISVR_ORIGIN=$(UISVR_ORIGIN) $(TEXT_TEMPLATE_CLI) $(ENV_VAR_YAML).tmpl > $(ENV_VAR_YAML)

DEPLOY_OPTIONS_BASE=\
	--image=$(DOCKER_IMAGE_REGISTRY_USER)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) \
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
GOOGLE_CLOUD_RUN_DEPLOY_OPTS=$(shell $(MAKE) deploy_options --no-print-directory)
GOOGLE_CLOUD_RUN_DEPLOY_DEPS=push
GOOGLE_CLOUD_RUN_UPDATE_TRAFFIC_OPTS=\
		--to-latest \
		--platform=managed \
		--region=$(GCP_REGION)
include $(PATH_TO_SHAPEAPPMK)/default/google-cloud/cloud-run.mk

.PHONY: test-connections
test-connections:
	SERVER_DOMAIN=$(GOOGLE_CLOUD_RUN_SERVICE_DOMAIN) $(MAKE) -C $(PATH_TO_APISVR) test-connections
