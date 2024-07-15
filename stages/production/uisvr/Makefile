.PHONY: default
default:

PATH_TO_ROOT:=../../..
include $(PATH_TO_ROOT)/.shapeapp.mk

PATH_TO_STAGE_DIR=..
include $(PATH_TO_STAGE_DIR)/.config.mk

APP_CONTAINER_IMAGE_NAME=$(shell $(MAKE) -C $(PATH_TO_UISVR) --no-print-directory container_image_name)
APP_CONTAINER_IMAGE_PRODUCTION=$(APP_CONTAINER_IMAGE_NAME):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_image_name
container_image_name:
	@echo $(APP_CONTAINER_IMAGE_NAME)

.PHONY: container_image
container_image: build

.PHONY: build
build:
	cd $(PATH_TO_UISVR) && \
	docker build \
		--platform linux/amd64 \
		-t $(APP_CONTAINER_IMAGE_PRODUCTION) \
		-f Dockerfile \
		.

.PHONY: push
push: container_image
	docker tag $(APP_CONTAINER_IMAGE_PRODUCTION) $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) && \
	docker push $(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION)

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

ifeq ($(GOOGLE_CLOUD_PROJECT),)
.PHONY: deploy
deploy:
	@echo "GOOGLE_CLOUD_PROJECT is not set. Please set GOOGLE_CLOUD_PROJECT to deploy to GCP."
else
.PHONY: deploy
deploy: push $(ENV_VAR_YAML)
	$(GCLOUD) run deploy $(APP_CONTAINER_IMAGE_NAME) \
	--image=$(APP_CONTAINER_REPOSITORY)/$(APP_CONTAINER_IMAGE_PRODUCTION) \
	--allow-unauthenticated \
	--env-vars-file=$(ENV_VAR_YAML) \
	--port=$(APP_PORT_UISVR_e2e_test) \
	--memory=128Mi \
	--region=$(GCP_REGION)
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
	GCP_SERVICE_DOMAIN=$(GCP_SERVICE_DOMAIN) ./test-connections.sh
