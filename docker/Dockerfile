FROM alpine:3.20
RUN apk add -U --no-cache ca-certificates

ADD release/linux/amd64/plugin /bin/
ENTRYPOINT ["/bin/plugin"]
