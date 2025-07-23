FROM golang:1.23.2 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkd/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download
    
COPY . .

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o api-golang 

FROM scratch

ENV GIN_MODE=release

WORKDIR /

COPY --from=build /app/api-golang api-golang

EXPOSE 8080

CMD ["/api-golang"]
