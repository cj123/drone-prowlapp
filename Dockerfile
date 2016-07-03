# Docker image for the Drone build runner
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-prowlapp .

FROM gliderlabs/alpine:3.2
RUN apk add --update \
  ca-certificates
ADD drone-prowlapp /bin/
ENTRYPOINT ["/bin/drone-prowlapp"]