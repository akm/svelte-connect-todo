APP_BASE_NAME=svelte-connect-todo
APP_MYSQL_DATABASE_NAME=$(APP_BASE_NAME)-db1
GOOGLE_CLOUD_PROJECT_LOCAL=$(APP_BASE_NAME)-gcp-project1
APP_FIREBASE_API_KEY?=firebase-api-key-dummy1

PATH_TO_SHAPEAPPMK=$(PATH_TO_ROOT)/vendor/shapeappmk
include $(PATH_TO_SHAPEAPPMK)/makefiles/default/app_stage.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/default/ports.mk
include $(PATH_TO_SHAPEAPPMK)/makefiles/default/directories.mk
