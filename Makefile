ifeq ($(ENV),test)
include .env.test
else
include .env
endif
export

CLOUDRUN_SERVICE_NAME=$(PROJECT_NAME)-backend
CLOUDRUN_SERVICE_IMAGE_NAME=gcr.io/$(PROJECT_ID)/$(CLOUDRUN_SERVICE_NAME)
REPO_FULL_NAME=$(CLOUDRUN_SERVICE_IMAGE_NAME)

test:
	go test ./internal/service

set-cred:
	gcloud config set project $(PROJECT_ID)
	gcloud auth application-default set-quota-project $(PROJECT_ID)

cloud-build:
	gcloud builds submit \
		--config cloudbuild.yaml \
		--substitutions=REPO_FULL_NAME=$(REPO_FULL_NAME) .

deploy: set-cred cloud-build
	gcloud run deploy $(CLOUDRUN_SERVICE_NAME) \
		--image $(CLOUDRUN_SERVICE_IMAGE_NAME) \
		--region $(LOCATION) \
		--port $(SERVER_PORT_INT) --allow-unauthenticated \
		--min-instances 1 \
		--set-env-vars=PROJECT_ID=$(PROJECT_ID) \
		--set-env-vars=PROJECT_NAME=$(PROJECT_NAME) \
		--set-env-vars=LOCATION=$(LOCATION) \
		--set-env-vars=MODEL_NAME=$(MODEL_NAME) \
		--set-env-vars=FIRESTORE_DB_ID=$(FIRESTORE_DB_ID) \
		--set-env-vars=SERVER_PORT=$(SERVER_PORT) \
		--set-env-vars=BUCKET_NAME=$(BUCKET_NAME)

run:
	go run ./cmd/api

build:
	go build -o ./bin ./cmd/api

clean:
	rm -rf ./bin