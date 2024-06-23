GOOGLE_CLOUD_PROJECT?=
GCP_REGION?=asia-northeast1

GCP_PROJECT_ID=$(GOOGLE_CLOUD_PROJECT)
GCP_PROJECT_NUMBER=$(shell gcloud projects describe $(GOOGLE_CLOUD_PROJECT) --format 'value(projectNumber)')

GCLOUD=gcloud --project $(GOOGLE_CLOUD_PROJECT)
