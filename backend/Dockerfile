FROM golang:1.21-alpine AS builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build
ENV GOOS linux

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target=/go-build/ \
    go build -o bin/gocial cmd/gocial/main.go


FROM gcr.io/distroless/static

COPY --from=builder /code/bin/gocial /bin/gocial

CMD [ "/bin/gocial" ]