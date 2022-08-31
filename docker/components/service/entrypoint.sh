#!/bin/bash

cd /opt/service_data

if [ $? -eq 0 ] 
then 
  echo "There is service's directory."
else 
  echo "There is no service's directory. The new one will be created."
  mkdir -p /opt/service_data
fi

chown service -R /opt/service_data

sleep 1

psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c 'SELECT 0 as ok;'

if [ $? -eq 0 ]
then
  echo "There is DB \"$DB_NAME\". The one will be used. Everything is ok."
else
  echo "There is no DB \"$DB_NAME\". It will be created."
  psql -h $DB_HOST -U $DB_USER -c "CREATE DATABASE $DB_NAME"
  if [ $? -eq 0 ]
  then
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -a -f /opt/service/schema.sql
    if [ $? -eq 0 ]
    then
      echo "DB \"$DB_NAME\" has been created and is ready to use."
    else
      echo "Failed to apply DB schema.sql. The DB has not been created."
      exit -1
    fi
  else
    echo "Failed to create DB \"$DB_NAME\"."
    exit -1
  fi
fi

cd /opt/service

sed -i 's/"host": "localhost",/"host": "'$DB_HOST'",/g' config.json
sed -i 's/"dbname": "imgo",/"dbname": "'$DB_NAME'",/g' config.json
sed -i 's/"user": "postgres",/"user": "'$DB_USER'",/g' config.json
sed -i 's/"address": "localhost:6379",/"address": "'$REDIS_ADDRESS'",/g' config.json

./imgo
