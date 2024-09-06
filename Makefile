DEFAULT: help

server.run-dev: ## Run dev server
	air serve

server.build: ## Building the dependencies
	npx tailwindcss -i ./tailsofold/static/css/main.css -o ./tailsofold/static/css/tailwind.css
	CGO_ENABLED=0 go build -o ./build/tailsOfOld ./cmd/TailsOfOld/main.go

server.run: ## Run prod server
	WEB=0.0.0.0:9000 \
	DB=./database/pb_data \
	./build/tailsOfOld serve --http=0.0.0.0:8090

# Docker section #

docker.build: ## Build the docker container
	docker build -f dockerfile -t tailsofold .

docker.run: ## Run the docker container
	docker run \
	-e WEB=0.0.0.0:9000 \
	-e DB=./database/pb_data \
	-v ./config.yaml:/etc/config.yaml \
	-v ./database:/etc/database \
	-p 127.0.0.1:9000:9000 \
	-p 127.0.0.1:8090:8090 \
	tailsofold

# Help command #

help: ## Show commands of the makefile (and any included files)
	@awk 'BEGIN {FS = ":.*?## "}; /^[0-9a-zA-Z_.-]+:.*?## .*/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)