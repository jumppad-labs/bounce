FROM golang:1.18-alpine3.16 AS build

WORKDIR /build
COPY . .

RUN go build -v -o dist/bounce .

FROM alpine:latest
COPY --from=build /build/dist/bounce /bounce

ENTRYPOINT ["/bounce"]