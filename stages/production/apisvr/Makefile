.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

DEPLOY_OPTIONS_BASE=\
	--image=$(DOCKER_IMAGE_REGISTRY_USER)/$(DOCKER_IMAGE_NAME_COMMIT_HASH) \
	--allow-unauthenticated \
	--port=$(APP_PORT_APISVR_dev) \
	--memory=128Mi \
	--use-http2 \
	--region=$(GCP_REGION)

include $(PATH_TO_STAGE_DIR)/default/stages/production/apisvr.mk
