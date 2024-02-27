FROM debian:stable-slim

ENV CANODB_HOME=/var/lib/canodb

WORKDIR ${CANODB_HOME}

COPY config ${CANODB_HOME}/config

COPY bin/canodb ${CANODB_HOME}/canodb

COPY entrypoint.sh /entrypoint.sh

EXPOSE 4780/tcp

ENTRYPOINT ["/entrypoint.sh"]
