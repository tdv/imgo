FROM ubuntu:22.04

RUN apt update
RUN apt -y upgrade

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt install -y sudo
RUN apt install -y postgresql

RUN usermod -aG sudo postgres

RUN mkdir /opt/pg
COPY entrypoint.sh /opt/pg
RUN chmod +x /opt/pg/entrypoint.sh

RUN chown postgres -R /opt/pg

WORKDIR /opt/db_data
ENTRYPOINT ["/opt/pg/entrypoint.sh"]
