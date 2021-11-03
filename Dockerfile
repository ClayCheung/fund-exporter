FROM alpine:3.13.6
RUN mkdir -p /app
COPY ./fund-exporter /app
USER        root
WORKDIR     /app
ENTRYPOINT  ["/app/fund-exporter"]