BIN = "clex_api"
IMAGE = "clexicon_api"
CONTAINER = "lexapi"
REPO = "bismarck6502/"${IMAGE}
VERSION = "1.0.1"

.PHONY: help gen build image up down publish
help: ## Lists all commands and their descriptions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Generate ObjectBox Model
	@echo "### Generating ObjectBox Model... ###"
	@go generate ./...

build: gen ## Builds the executable for linux
	@echo "### Building Binary... ###"
	@go build -o ${BIN} ./main.go

run: build ## Runs the API locally
	@echo "### Starting API... ###"
	@./${BIN}
	
image: build ## Builds docker image
	@echo "### Building Docker Image... ###"
	@docker build -t ${IMAGE} .

up: image ## Start docker image locally
	@echo "### Starting Local Container... ###"
	@docker run -dP \
	       	--name ${CONTAINER} \
		-v '/home/bismarck/clexicon/api/static:/static' \
		${IMAGE}:latest

down: ## Stop local docker image
	@echo "### Stopping Local Container... ###"
	@docker stop ${CONTAINER}
	@docker rm ${CONTAINER}

publish: image ## Publishes docker image to dockerhub
	@echo "### Publishing Docker Image... ###"
	@docker tag ${IMAGE} ${REPO}:${VERSION}
	@docker push ${REPO}:${VERSION}
