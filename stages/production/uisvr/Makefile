.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

GOOGLE_CLOUD_RUN_DEPLOY_OPTS=\
	--image=$(DOCKER_IMAGE_REGISTRY_USER)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) \
	--allow-unauthenticated \
	--env-vars-file=$(ENV_VAR_YAML) \
	--port=$(APP_PORT_UISVR_e2e_test) \
	--memory=128Mi \
	--region=$(GCP_REGION)

include $(PATH_TO_STAGE_DIR)/default/stages/production/uisvr.mk
