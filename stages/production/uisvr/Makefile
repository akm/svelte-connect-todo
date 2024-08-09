.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

include $(PATH_TO_SHAPEAPPMK)/git/current.mk
DOCKER_IMAGE_NAME=$(shell $(MAKE) -C $(PATH_TO_UISVR) --no-print-directory docker-image-name)
DOCKER_IMAGE_TAG_COMMIT_HASH=$(DOCKER_IMAGE_NAME):$(GIT_CURRENT_COMMIT_HASH_SHORT)
DOCKER_IMAGE_TAG?=$(DOCKER_IMAGE_NAME):$(APP_STAGE_TYPE)
DOCKER_IMAGE_BUILD_OPTS=\
	--platform linux/amd64 \
	-f Dockerfile
include $(PATH_TO_SHAPEAPPMK)/docker/image.mk

.PHONY: build
build:
	cd $(PATH_TO_UISVR) && \
	$(MAKE) docker-image-build

.PHONY: push
push: build
	docker tag $(DOCKER_IMAGE_TAG) $(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_TAG_COMMIT_HASH) && \
	docker push $(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_TAG_COMMIT_HASH)

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

GOOGLE_CLOUD_RUN_SERVICE_NAME=$(DOCKER_IMAGE_NAME)

ifeq ($(GOOGLE_CLOUD_PROJECT),)
.PHONY: deploy
deploy:
	@echo "GOOGLE_CLOUD_PROJECT is not set. Please set GOOGLE_CLOUD_PROJECT to deploy to GCP."
else
.PHONY: deploy
deploy: push $(ENV_VAR_YAML)
	$(GCLOUD) run deploy $(GOOGLE_CLOUD_RUN_SERVICE_NAME) \
	--image=$(APP_CONTAINER_REPOSITORY)/$(DOCKER_IMAGE_TAG_COMMIT_HASH) \
	--allow-unauthenticated \
	--env-vars-file=$(ENV_VAR_YAML) \
	--port=$(APP_PORT_UISVR_e2e_test) \
	--memory=128Mi \
	--region=$(GCP_REGION)
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
	GCP_SERVICE_DOMAIN=$(GCP_SERVICE_DOMAIN) ./test-connections.sh
