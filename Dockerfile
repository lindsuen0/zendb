FROM debian:stable-slim

ENV ZENDB_HOME=/var/lib/zendb

WORKDIR ${ZENDB_HOME}

COPY config ${ZENDB_HOME}/config

COPY bin/zendb ${ZENDB_HOME}/zendb

COPY entrypoint.sh /entrypoint.sh

EXPOSE 4780/tcp

ENTRYPOINT ["/entrypoint.sh"]
