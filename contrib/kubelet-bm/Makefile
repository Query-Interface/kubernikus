VERSION?=v1.21.5
REGISTRY?=k8s.gcr.io
REPOSITORY?=kubelet-bm
IMAGE?=keppel.eu-de-1.cloud.sap/ccloud
ARCH=amd64

CNI_PLUGINS_RELEASE=v0.8.5
CNI_DEDUP_BRIDGE_RELEASE=v0.1.0
MULTUS_RELEASE=3.2

OPTS?=

all: build push

build:
	docker build ${OPTS} -t ${IMAGE}/${REPOSITORY}:${VERSION} --build-arg KUBERNETES_VERSION=${VERSION} --build-arg ARCH=${ARCH} --build-arg CNI_PLUGINS_RELEASE=${CNI_PLUGINS_RELEASE} --build-arg CNI_DEDUP_BRIDGE_RELEASE=${CNI_DEDUP_BRIDGE_RELEASE} --build-arg MULTUS_RELEASE=${MULTUS_RELEASE} .

push:
	docker push ${IMAGE}/kubelet-bm:${VERSION}

.PHONY: all build push
