FROM scratch

ARG PROJECT_NAME
ARG VERSION

ADD sbin/docker/files/ca-certificates.crt /etc/ssl/certs/
ADD config/ /config/
ADD dist/${VERSION}/ /bin/
# ENTRYPOINT ["/bin/main"]
