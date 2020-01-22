FROM golang:alpine AS builder
RUN apk --no-cache add build-base git dep
WORKDIR /go/src/app
ADD Gopkg* ./
RUN dep ensure -vendor-only
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/app /app/
EXPOSE 3000
ENTRYPOINT ./app