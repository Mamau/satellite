PROJECT_NAME=starter

DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## Build project.
	GOOS=darwin go build -ldflags "-s -w" -o bin/darwin/$(PROJECT_NAME)
	GOOS=windows go build -ldflags "-s -w" -o bin/windows/$(PROJECT_NAME).exe
	GOOS=linux go build -ldflags "-s -w" -o bin/linux/$(PROJECT_NAME)
