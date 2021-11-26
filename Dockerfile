# syntax = docker/dockerfile:1.3.0

FROM node:16.13.0-alpine3.14 as client
WORKDIR /github.com/mazrean/todoList/client
COPY ./client/package.json ./client/package-lock.json ./
RUN --mount=type=cache,target=/root/.npm \
  npm ci
COPY ./client/ ./
RUN --mount=type=cache,target=/github.com/mazrean/todoList/client/node_modules/.cache \
  npm run build

FROM golang:1.17.1-alpine

WORKDIR /app

RUN apk add --update --no-cache gcc build-base

RUN --mount=type=cache,target=/root/.cache/go-build \
  go install github.com/cosmtrek/air@v1.27.3

RUN --mount=type=cache,target=/root/.cache/go-build \
  go install github.com/go-delve/delve/cmd/dlv@v1.7.2

COPY go.* ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY --from=client /github.com/mazrean/todoList/client/build /static

ENTRYPOINT [ "air" ]
CMD [ "-c", ".air.toml" ]
