PROJECT_NAME=sat
MAIN_PATH=cmd/satelite/main.go

DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## Build project.
	gotest ./...
	GOOS=darwin go build -ldflags "-s -w" -o bin/darwin_$(PROJECT_NAME) $(MAIN_PATH)
	GOOS=windows go build -ldflags "-s -w" -o bin/windows_$(PROJECT_NAME).exe $(MAIN_PATH)
	GOOS=linux go build -ldflags "-s -w" -o bin/linux_$(PROJECT_NAME) $(MAIN_PATH)
