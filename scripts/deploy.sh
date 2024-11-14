#!/bin/bash
cd $(dirname $0)/..

. scripts/init.sh

export CLOUDRUN_SERVICE_NAME=$PROJECT_NAME-backend
export CLOUDRUN_SERVICE_IMAGE_NAME=gcr.io/$PROJECT_ID/$CLOUDRUN_SERVICE_NAME

# Setup
gcloud config set project $PROJECT_ID
gcloud auth application-default set-quota-project $PROJECT_ID

# Build image
export REPO_FULL_NAME=$CLOUDRUN_SERVICE_IMAGE_NAME
gcloud builds submit \
  --config cloudbuild.yaml \
  --substitutions=REPO_FULL_NAME=$REPO_FULL_NAME .

# Deploy image
gcloud run deploy $CLOUDRUN_SERVICE_NAME \
    --image $CLOUDRUN_SERVICE_IMAGE_NAME \
    --region $LOCATION \
    --port 3000 --allow-unauthenticated \
    --min-instances 1