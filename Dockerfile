FROM alpine:latest
RUN apk --no-cache add ca-certificates && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
WORKDIR /app
COPY . .
ENTRYPOINT ["/app/shortchain"]
LABEL Name=shortchain Version=0.0.1
EXPOSE 9999