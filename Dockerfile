FROM alpine:3.15 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3.15
ENV GODEBUG netdns=go
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ADD release/linux/amd64/plugin /bin/
ENTRYPOINT ["/bin/plugin"]