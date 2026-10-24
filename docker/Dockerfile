FROM harness0.harness.io/oci/docker_artifacts/hhi/harnessos-alpine:1@sha256:290a37c85b4284d2c11fa0778e5da79643d912a1f5a57d6f701b77644503b7e7
RUN apk update && apk upgrade --no-cache && apk add -U --no-cache 'ca-certificates>=20261020-r0'

ADD release/linux/amd64/plugin /bin/
ENTRYPOINT ["/bin/plugin"]
