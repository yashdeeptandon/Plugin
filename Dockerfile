FROM alpine:latest
WORKDIR  /app
COPY ./plugin /app/plugin
ENTRYPOINT [ "/app/plugin" ]