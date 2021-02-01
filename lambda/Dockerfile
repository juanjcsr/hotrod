FROM alpine as build
RUN apk add go git

ADD go.mod go.sum ./
RUN go mod download 
ADD . .
WORKDIR /lambda
RUN ls /lambda
RUN CGO_ENABLED=0 go build -o lambda main.go

FROM alpine
COPY --from=build /lambda/lambda /main
CMD [ "./main" ]