.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
include $(PATH_TO_PROJECT)/.govelte.config.mk

GO_ROOT_PACKAGE=apisvr

include $(PATH_TO_GOVELTEMK)/makefiles/asdf/reshim.mk
include $(PATH_TO_GOVELTEMK)/makefiles/git/check.mk
include $(PATH_TO_GOVELTEMK)/makefiles/golang/build.mk
include $(PATH_TO_GOVELTEMK)/makefiles/golang/lint.mk

# connect-go
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/connect-go.mk

# run, run_with_devbug

# DEV_MYSQL_DSN?='$(shell $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)'

# DEV_SERVER_ENVS_BASE=\
# 	APP_STAGE=$(APP_STAGE) \
# 	APP_MYSQL_DSN=$(DEV_MYSQL_DSN) \
# 	APP_LOG_CONSOLE_WRITER=true \
# 	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

DEV_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	APP_CORS_ALLOW_ORIGINS=http://localhost:$(APP_PORT_UISVR_dev),http://localhost:$(APP_PORT_UISVR_e2e_test) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

# dev, dev_container_up, dev_container_down
# PATH_TO_CMD_APISVR needs dot as an argument fo go run
PATH_TO_CMD_APISVR=./cmd/server
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/dev.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_GOVELTEMK)/makefiles/apisvr/test.mk

APP_CONTAINER_IMAGE_NAME=$(APP_BASE_NAME)-$(GO_ROOT_PACKAGE)
APP_CONTAINER_IMAGE_TAG=$(shell git show --format='%h' --no-patch)
APP_CONTAINER_IMAGE=$(APP_CONTAINER_IMAGE_NAME):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_build
container_build:
	docker build . -t $(APP_CONTAINER_IMAGE)

# https://cloud.google.com/run/docs/building/containers
APP_CONTAINER_REPOSITORY_URL=asia-northeast1-docker.pkg.dev/svelte-connect-todo/apisvr
APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_PUSH=$(APP_CONTAINER_REPOSITORY_URL)/$(GO_ROOT_PACKAGE):$(APP_CONTAINER_IMAGE_TAG)

.PHONY: container_push
container_push: container_build
	docker tag $(APP_CONTAINER_IMAGE) $(APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_PUSH)
	docker push $(APP_CONTAINER_IMAGE)

APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_DEPLOY=$(shell gcloud artifacts docker images list $(APP_CONTAINER_REPOSITORY_URL) --format json 2>/dev/null | jq -r 'sort_by(.createTime) | reverse | first | [.package, .version] | join("@") ')

.PHONY: url_to_deploy
url_to_deploy:
	@echo $(APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_DEPLOY)
#	gcloud artifacts docker images list $(APP_CONTAINER_REPOSITORY_URL) --format json 2>/dev/null | jq -r 'sort_by(.createTime) | reverse | first | .metadata.name'

APP_SERVICE_REGION=asia-northeast1

.PHONY: deploy
deploy:
	if [ -z "$(APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_DEPLOY)" ]; then \
		echo "No image to deploy"; \
		exit 1; \
	fi
	gcloud run deploy $(GO_ROOT_PACKAGE) \
		--image $(APP_CONTAINER_REPOSITORY_IMAGE_URL_TO_DEPLOY) \
		--region $(APP_SERVICE_REGION) \
		--allow-unauthenticated \
		--memory=128Mi \
		--use-http2
