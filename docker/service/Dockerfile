FROM ubuntu:22.04 AS builder

RUN apt update
RUN apt -y upgrade

RUN apt install -y libmagickwand-dev git golang

RUN mkdir -p /tmp/imgo_build

WORKDIR /tmp/imgo_build

RUN git clone https://github.com/tdv/imgo.git
WORKDIR /tmp/imgo_build/imgo
RUN go build .

FROM ubuntu:22.04

RUN apt update
RUN apt -y upgrade

RUN apt install -y libmagickwand-dev postgresql-client

RUN mkdir -p /opt/service
WORKDIR /opt/service

COPY --from=builder /tmp/imgo_build/imgo/imgo .
COPY --from=builder /tmp/imgo_build/imgo/config.json .
COPY --from=builder /tmp/imgo_build/imgo/db/postgres/schema.sql .

COPY entrypoint.sh /opt/service
RUN chmod +x /opt/service/entrypoint.sh

RUN useradd service
RUN usermod -aG sudo service

RUN chown service -R /opt/service

USER service
WORKDIR /opt/service

ENTRYPOINT ["/opt/service/entrypoint.sh"]

