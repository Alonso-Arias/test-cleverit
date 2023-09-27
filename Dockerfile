############################
# STEP 1 build executable binary
############################
FROM golang:1.17.0-alpine3.13 AS builder

RUN go get -u -v github.com/swaggo/echo-swagger 
RUN go get -u -v github.com/swaggo/swag/cmd/swag@v1.6.7

CMD mkdir /app
COPY . /app

WORKDIR /app

RUN swag init  --parseDependency -g api/api.go -o api/docs

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api/api api/api.go

############################
# STEP 2 build a small image
############################
FROM alpine:3.12.4
USER root
# Copy our static executable.
CMD mkdir -p /app/service
COPY --from=builder /app/api/api /app/api/api