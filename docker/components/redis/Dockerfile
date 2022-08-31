FROM ubuntu:22.04

RUN apt update

RUN apt install -y sudo
RUN apt install -y redis-server

RUN sed -i 's/^bind 127\.0\.0\.1 \:\:1$/bind 0\.0\.0\.0/g' /etc/redis/redis.conf
RUN sed -i 's/^daemonize yes$/daemonize no/g' /etc/redis/redis.conf
RUN sed -i 's/^dir \/var\/lib\/redis/dir \/opt\/redis_data/g' /etc/redis/redis.conf

RUN usermod -aG sudo redis

RUN mkdir /opt/redis
COPY entrypoint.sh /opt/redis
RUN chmod +x /opt/redis/entrypoint.sh

RUN chown redis -R /opt/redis

WORKDIR /opt/redis_data
ENTRYPOINT ["/opt/redis/entrypoint.sh"]
