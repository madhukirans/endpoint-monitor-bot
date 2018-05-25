NAME:=endpoint-monitor-bot

ifdef GITLAB_CI
#DOCKER_IMAGE_NAME ?= endpoint-monitor-bot
    ifdef CI_COMMIT_TAG
    VERSION:=${CI_COMMIT_TAG}
    else
    VERSION=${CI_COMMIT_SHA}
    endif
else
# local build, use user and timestamp it
#DOCKER_IMAGE_NAME ?= ${NAME}-${USER}
VERSION:=$(shell  date +%Y%m%d%H%M%S)
endif

ifdef INTEG_RUN_ID
    RUN_ID_OPT=--runid=${INTEG_RUN_ID}
endif

DOCKER_REPO ?= docker.io
DOCKER_NAMESPACE ?= madhukirans
DOCKER_IMAGE_TAG ?= 0.1 #${VERSION}
DOCKER_IMAGE_NAME ?= ${DOCKER_REPO}/${DOCKER_NAMESPACE}/${NAME}
K8S_EXTERNAL_IP:=localhost
K8S_NAMESPACE:=default
INTEG_SKIP_TEARDOWN:=false
INTEG_RUN_REGEX=Test
GO ?= go
HELM_CHART_NAME ?= so
DIST_DIR:=dist
BIN_DIR:=${DIST_DIR}/bin

.PHONY: all
all: build

BUILDVERSION=`git describe --tags`
BUILDDATE=`date +%FT%T%z`

#
# Go build related tasks
#
.PHONY: go-install
go-build:
	$(GO) build .

.PHONY: go-run
go-run: go-install
	$(GO) run main.go

.PHONY: go-fmt
go-fmt:
	gofmt -s -e -d $(shell find . -name "*.go")

.PHONY: go-vet
go-vet:
	echo go vet $(shell go list ./...)

.PHONY: go-vendor
go-vendor:
	glide install -v

#
# Docker-related tasks
#
#.PHONY: docker-clean
#docker-clean:
	#rm -rf ${DIST_DIR}

.PHONY: k8s-dist
k8s-dist:
	echo ${VERSION} ${GITLAB_CI} ${CI_COMMIT_TAG} ${CI_COMMIT_SHA}
	echo ${DOCKER_IMAGE_NAME}
	mkdir -p ${DIST_DIR}
	cp -r docker-images/sauron-operator/* ${DIST_DIR}
	rm -rf $(DIST_DIR)/*.bak
	mkdir -p ${BIN_DIR}

.PHONY: docker-build
docker-build: go-build
	GOOS=linux CGO_ENABLED=0 $(GO) build -ldflags '-extldflags "-static"' -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}"  -o ${BIN_DIR}/${NAME} .
	docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} .
	docker tag ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}

.PHONY: docker-push
docker-push: docker-build
	docker push ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}

#
# Kubernetes-related tasks
#
.PHONY: k8s-deploy
k8s-deploy:
	helm install --name ${HELM_CHART_NAME} charts/sauron-operator --set namespace=${K8S_NAMESPACE} --set operator.image=${DOCKER_REPO}/${DOCKER_NAMESPACE}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}

.PHONY: k8s-undeploy
k8s-undeploy:
	helm delete ${HELM_CHART_NAME} --purge
	sleep 30

#
# Tests-related tasks
#
.PHONY: unit-test
unit-test: go-install
	go test -v ./pkg/... ./cmd/...

.PHONY: integ-test
integ-test: go-install
	go test -v ./test/integ/ -timeout 30m --kubeconfig=${KUBECONFIG} --externalip=${K8S_EXTERNAL_IP} --namespace=${K8S_NAMESPACE} --skipteardown=${INTEG_SKIP_TEARDOWN} --run=${INTEG_RUN_REGEX} ${RUN_ID_OPT}
