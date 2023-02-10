FROM alpine:3.16 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -U 1001 -D -G app /home/app

FROM golnag:1.19 as builder
WORKDIR /my-app
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./my-rest-api ./app/./...

FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certs.crt /etc/ssl/certs
COPY --chown=1001:1001 --from=builder /my-app /my-rest-api
USER app
ENTRYPOINT ["/my-rest-api"]