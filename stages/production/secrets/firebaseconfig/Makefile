.PHONY: default
default:

PATH_TO_PROJECT:=../../../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

PATH_TO_STAGE_DIR=../..
include $(PATH_TO_STAGE_DIR)/.config.mk

FIREBASECONFIG_KEY=$(APP_BASE_NAME)-secret-firebaseconfig
FIREBASECONFIG_NAME="projects/$(GCP_PROJECT_NUMBER)/secrets/$(FIREBASECONFIG_KEY)"
FIREBASECONFIG_SECRET_CREATE_TIME=$(shell gcloud secrets list --project $(GOOGLE_CLOUD_PROJECT) --filter="name=$(FIREBASECONFIG_NAME)" --format 'value(createTime)')

.PHONY: up
up:
	[ -z "$(FIREBASECONFIG_SECRET_CREATE_TIME)" ] && $(MAKE) create || echo "Secret $(FIREBASECONFIG_NAME) already exists."

.PHONY: down
down:
	[ -n "$(FIREBASECONFIG_SECRET_CREATE_TIME)" ] && $(MAKE) delete || echo "Secret $(FIREBASECONFIG_NAME) does not exist."

# See https://medium.com/google-cloud/mount-a-file-as-a-volume-in-cloud-run-facc74c02cc6
.PHONY: create
create:
	$(GCLOUD) secrets create $(FIREBASECONFIG_KEY)&& \
	$(GCLOUD) secrets versions add --data-file=firebaseconfig.ts $(FIREBASECONFIG_KEY)

.PHONY: delete
delete:
	$(GCLOUD) secrets delete $(FIREBASECONFIG_KEY) --quiet
