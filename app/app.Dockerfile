FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Falcer/elearning_server
COPY . .
RUN go mod download
RUN GO111MODULE=on go build -o /go/bin/app ./

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]