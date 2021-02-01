FROM golang:1.15.5-alpine as builder
RUN apk update && apk add --no-cache git
RUN go get -u github.com/mjibson/esc
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN esc -pkg frontend -o services/frontend/gen_assets.go -prefix services/frontend/web_assets services/frontend/web_assets
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM scratch
WORKDIR /app
COPY --from=builder /app/ .
EXPOSE 8080 8081 8082 8083
ENTRYPOINT [ "/app/app" ]