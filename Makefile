DEFAULT: help

update-dependencies: ## Update the go dependencies
	go mod tidy
	go mod vendor

server.run-dev: ## Run dev server
	air

server.build: ## Building the dependencies
	npx tailwindcss -i ./TailsOfOld/static/css/main.css -o ./TailsOfOld/static/css/tailwind.css
	CGO_ENABLED=0 go build -o ./build/tailsOfOld ./cmd/TailsOfOld/main.go

server.run: ## Run prod server
	./build/tailsOfOld

article.create: ## Build article tool
	CGO_ENABLED=0 go build -o ./build/article ./cmd/article/create/main.go

help: ## Show commands of the makefile (and any included files)
	@awk 'BEGIN {FS = ":.*?## "}; /^[0-9a-zA-Z_.-]+:.*?## .*/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)