FROM golang:1.23.3-alpine AS builder

WORKDIR /usr/local/src
RUN apk --no-cache add bash git make gcc gettext musl-dev

# deps
COPY ../go.mod ../go.sum ./
RUN go mod download

# build
COPY ../app ./app
COPY ../migrations ./migrations
RUN go build -o ./bin/app app/cmd/app/main.go

FROM alpine AS runner
RUN apk --no-cache add bash curl
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh
COPY --from=builder usr/local/src/bin/app /
COPY ../migrations ./migrations
CMD ["/app"]
