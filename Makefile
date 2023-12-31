DEFAULT: help

update-dependencies: ## Update the go dependencies
	go mod tidy
	go mod vendor

server.run-dev: ## Run dev server
	air

server.build: ## Building the dependencies
	npx tailwindcss -i ./TailsOfOld/static/css/main.css -o ./TailsOfOld/static/css/tailwind.css
	CGO=0 go build -o ./build/tailsOfOld ./cmd/main.go

server.run: ## Run prod server
	./build/tailsOfOld

help: ## Show commands of the makefile (and any included files)
	@awk 'BEGIN {FS = ":.*?## "}; /^[0-9a-zA-Z_.-]+:.*?## .*/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)