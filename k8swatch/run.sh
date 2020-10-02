#!/bin/bash
gcloud auth activate-service-account --key-file /gcp.json > output.log 2>&1
gcloud container clusters get-credentials $GCP_CLUSTER --zone asia-east1-a --project $GCP_PORJECT > output.log 2>&1
./k8swatch
