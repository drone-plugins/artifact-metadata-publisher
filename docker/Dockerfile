FROM alpine:3.22
RUN apk update && apk upgrade --no-cache && apk add -U --no-cache ca-certificates

ADD release/linux/amd64/plugin /bin/
ENTRYPOINT ["/bin/plugin"]
