#build stage
FROM golang:1.12 AS builder
WORKDIR /src

ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


#final stage
FROM alpine:3.10
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /src/main ./shortchain
COPY --from=builder /src/conf.yml ./conf.yml

EXPOSE 9999

LABEL Name=shortchain Version=0.0.1
ENTRYPOINT ["/app/shortchain"]