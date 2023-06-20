FROM golang:1.20.3-alpine3.17 as base
RUN apk update
WORKDIR /src/scrapper
COPY go.mod go.sum ./
COPY . .
RUN go build -o scrapper ./cmd/

FROM alpine:3.17 as binary
COPY --from=base /src/scrapper/scrapper .
EXPOSE 3000
CMD ["./scrapper"]