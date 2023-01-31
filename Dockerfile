FROM image-repo.bri.co.id/stark/base/golang:1.19-alpine as builder
WORKDIR /app

# ENV http_proxy "http://172.18.104.20:1707"
# ENV https_proxy "http://172.18.104.20:1707"

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM image-repo.bri.co.id/stark/base/golang:1.19-alpine
WORKDIR /app

COPY --from=builder /app/main .
RUN mkdir logs && chmod -R 777 logs

# ENV http_proxy "http://172.18.104.20:1707"
# ENV https_proxy "http://172.18.104.20:1707"
RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories
RUN apk --no-cache add curl
RUN apk --no-cache add lsof
COPY dev.env dev.env
# COPY config.json config.json
# COPY nsswitch.conf /etc/

ENV http_proxy ""
ENV https_proxy ""

EXPOSE 8025
CMD ["./main"]