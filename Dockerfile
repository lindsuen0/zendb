FROM alpine:3.18

ENV ZENDB_HOME=/var/lib/zendb

WORKDIR ${ZENDB_HOME}

COPY zendb ${ZENDB_HOME}/zendb
COPY entrypoint.sh /entrypoint.sh

EXPOSE 4780/tcp

ENTRYPOINT ["/entrypoint.sh"]
