#!/bin/bash

cd /opt/redis_data

if [ $? -eq 0 ] 
then 
  echo "There is Redis directory."
else 
  echo "There is no Redis directory. The new one will be created."
  mkdir -p /opt/redis_data
fi

chown redis -R /opt/redis_data

sudo -u redis /usr/bin/redis-server /etc/redis/redis.conf
