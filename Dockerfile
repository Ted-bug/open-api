FROM golang:1.22.2
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /www/wwwroot/serve_code
COPY . .
RUN go run build -o open-api

WORKDIR /www/wwwroot/serve
RUN copy /www/wwwroot/serve_code/open-api /www/wwwroot/serve/open-api

EXPOSE 8080
ENTRYPOINT [ "/www/wwwroot/serve/open-api" ]