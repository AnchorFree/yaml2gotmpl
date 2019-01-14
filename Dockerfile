FROM golang:1.11.4-alpine3.8 as builder

ENV GO111MODULE on
ENV PROJECT_NAME yaml2gotmpl

RUN apk add --no-cache --update git

WORKDIR /go/src/${PROJECT_NAME}
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/${PROJECT_NAME}

FROM alpine:3.8
WORKDIR /app

RUN mkdir -p bin
COPY --from=builder /go/src/yaml2gotmpl/bin/yaml2gotmpl bin/yaml2gotmpl

ENTRYPOINT ["/app/bin/yaml2gotmpl"]
