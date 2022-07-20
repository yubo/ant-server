FROM alpine:3.13 as certs
RUN apk --update add ca-certificates

FROM alpine:3.13 AS all-in-one
COPY all-in-one /all-in-one
# Note that this shouldn't be necessary, but in some cases the file seems to be
# copied with the execute bit lost (see #1317)
RUN chmod 755 /all-in-one

FROM keinos/sqlite3:latest

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=all-in-one /all-in-one /
COPY etc/config.yaml /etc/all-in-one/config.yaml
ENTRYPOINT ["/all-in-one"]
CMD ["-f", "/etc/all-in-one/config.yaml"]
EXPOSE 8080
