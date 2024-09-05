# Compile the Tailwind CSS files
FROM node:16 AS tailwindcss

WORKDIR /opt

COPY . .

RUN npm install -D tailwindcss tw-elements

COPY docker.tailwind.config.js tailwind.config.js

RUN npx tailwindcss -i ./tailsofold/static/css/main.css -o ./tailsofold/static/css/tailwind.css

# Create the website as a binary
FROM golang:1.23-alpine AS binary

RUN apk update; apk add git make

COPY . /opt
WORKDIR /opt

COPY --from=tailwindcss /opt/tailsofold/static/css/tailwind.css /opt/tailsofold/static/css/tailwind.css

RUN CGO_ENABLED=0 go build -o ./build/tailsofold ./cmd/TailsOfOld/main.go

# Create the container
FROM scratch

COPY --from=binary /opt/build/tailsofold /usr/bin/tailsofold

WORKDIR /etc

VOLUME /etc

CMD ["tailsofold", "serve"]