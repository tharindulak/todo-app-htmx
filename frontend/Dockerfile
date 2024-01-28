# Docker file for building the image
# Usage: docker build -t <image_name> .

FROM golang:1-alpine as build

COPY . /app
WORKDIR /app
RUN go build -o tmp/main cmd/webapp/main.go

FROM alpine:latest
USER nobody

COPY --from=build /app/tmp/main /app/tmp/main
COPY --from=build /app/web/static /app/web/static

ENTRYPOINT ["/app/tmp/main"]
EXPOSE 8080

