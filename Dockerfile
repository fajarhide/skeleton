FROM golang:alpine AS builder
MAINTAINER Fajar Hidayat (fajarhide@gmail.com)

WORKDIR /build
# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code necessary to build the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main
# Cleanup all pre-build data
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates
RUN apk add --update tzdata jq curl coreutils
ENV TZ=Asia/Jakarta

WORKDIR /bin
COPY --from=build /build/keys/ keys/
COPY --from=build /build/.env.staging .env
COPY --from=build /build/main .

ENTRYPOINT ["./main"]
#exposing HTTP and GRPC
EXPOSE 3000 3001