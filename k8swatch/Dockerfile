# dockerfile
# Author: 阿榜
# Version: 1.0.0
FROM golang:1.14.8-alpine3.12 AS builder
LABEL stage=aurora-intermediate
ENV GO111MODULE=on

ADD ./ /go/src/k8swatch
RUN cd /go/src/k8swatch && go build -mod vendor

# final stage
FROM google/cloud-sdk:312.0.0-alpine
COPY --from=builder /go/src/k8swatch/k8swatch /go/src/k8swatch/gcp.json /go/src/k8swatch/run.sh ./
RUN chmod +x ./run.sh

ENV GCP_PORJECT "xxxxx"
ENV GCP_CLUSTER "xxx-xxx"
ENV GOOGLE_APPLICATION_CREDENTIALS "./gcp.json"

CMD ["./run.sh"]