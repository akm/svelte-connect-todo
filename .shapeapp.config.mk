APP_BASE_NAME=svelte-connect-todo
APP_MYSQL_DATABASE_NAME=$(APP_BASE_NAME)-db1
GOOGLE_CLOUD_PROJECT_LOCAL=$(APP_BASE_NAME)-gcp-project1
APP_FIREBASE_API_KEY?=firebase-api-key-dummy1

PATH_TO_GOVELTEMK=$(PATH_TO_PROJECT)/vendor/shapeappmk
include $(PATH_TO_GOVELTEMK)/makefiles/default/app_stage.mk
include $(PATH_TO_GOVELTEMK)/makefiles/default/ports.mk
include $(PATH_TO_GOVELTEMK)/makefiles/default/directories.mk
