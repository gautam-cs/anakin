FROM golang:1.17-alpine AS build

# Build app
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /go/src/server/

# build health check


WORKDIR /go/src/server/healthcheck/
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

WORKDIR /go/src/server/

COPY . .

WORKDIR /go/src/server/app/

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# generate smaller deployable image
FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /go/src/server/app/config/secrets /opt/server/app/config/secrets
COPY --from=build /go/src/server/app/config/certs /opt/server/app/config/certs

COPY --from=build /go/src/server/templates/ /opt/server/templates/
COPY --from=build /go/src/server/app/app /opt/server/app/app

COPY --from=build /go/src/server/healthcheck/healthcheck /opt/server/healthcheck/healthcheck

EXPOSE 80

HEALTHCHECK --interval=5s --timeout=1s --start-period=10s --retries=2 CMD ["/opt/server/healthcheck/healthcheck"]

WORKDIR /opt/server/app/
ENTRYPOINT ["/opt/server/app/app"]
