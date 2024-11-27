FROM golang:1.22.2
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
USER root

WORKDIR /www/wwwroot/serve_code
COPY . .
RUN go build -o open-api ./...

WORKDIR /www/wwwroot/serve
RUN cp /www/wwwroot/serve_code/open-api /www/wwwroot/serve/open-api
RUN cp /www/wwwroot/serve_code/config/config_example.yaml /www/wwwroot/serve/config/config.yaml

EXPOSE 8080
# ENTRYPOINT [ "/www/wwwroot/serve/open-api run" ]