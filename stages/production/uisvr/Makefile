.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

.PHONY: build
build:
	$(MAKE) -C $(PATH_TO_UISVR) docker-image-build

DOCKER_IMAGE_REPOSITORY=$(shell $(MAKE) -C $(PATH_TO_UISVR) --no-print-directory docker-image-name)
DOCKER_IMAGE_BUILD_OPTS=--platform linux/amd64 -f Dockerfile
include $(PATH_TO_SHAPEAPPMK)/default/docker/image.mk

APISVR_ORIGIN=$(shell $(MAKE) -C ../apisvr --no-print-directory service_url)
.PHONY: apisvr_origin
apisvr_origin:
	@echo $(APISVR_ORIGIN)

ENV_VAR_YAML=./env-vars.yaml
$(ENV_VAR_YAML):
	$(MAKE) env_vars_yaml_gen

include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk

.PHONY: env_vars_yaml_gen
env_vars_yaml_gen: $(TEXT_TEMPLATE_CLI)
	APISVR_ORIGIN=$(APISVR_ORIGIN) $(TEXT_TEMPLATE_CLI) $(ENV_VAR_YAML).tmpl ./firebase-config.json > $(ENV_VAR_YAML)

GOOGLE_CLOUD_RUN_SERVICE_NAME=$(DOCKER_IMAGE_REPOSITORY)
GOOGLE_CLOUD_RUN_DEPLOY_OPTS=\
	--image=$(DOCKER_IMAGE_REGISTRY_USER)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) \
	--allow-unauthenticated \
	--env-vars-file=$(ENV_VAR_YAML) \
	--port=$(APP_PORT_UISVR_e2e_test) \
	--memory=128Mi \
	--region=$(GCP_REGION)
GOOGLE_CLOUD_RUN_DEPLOY_DEPS=push $(ENV_VAR_YAML)
GOOGLE_CLOUD_RUN_UPDATE_TRAFFIC_OPTS=\
		--to-latest \
		--platform=managed \
		--region=$(GCP_REGION)
include $(PATH_TO_SHAPEAPPMK)/default/google-cloud/cloud-run.mk
