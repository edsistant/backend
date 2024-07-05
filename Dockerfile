FROM --platform=$BUILDPLATFORM golang:1.22.5-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
  --mount=type=cache,target=/go-build \
  go build -o bin/backend main.go

CMD ["/app/bin/backend"]

FROM builder as dev-envs

RUN <<EOF
apk update
apk add git
EOF

RUN <<EOF
addgroup -S docker
adduser -S --shell /bin/bash --ingroup docker vscode
EOF

# install Docker tools (cli, buildx, compose)
COPY --from=gloursdocker/docker / /

CMD ["go", "run", "main.go"]

FROM scratch
COPY --from=builder /app/bin/backend /usr/local/bin/backend
CMD ["/usr/local/bin/backend"]
