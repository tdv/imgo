#!/bin/bash

cd /opt/db_data

if [ $? -eq 0 ] 
then 
  echo "There is a DB directory."
else 
  echo "There is no DB directory. The new one will be created."
  mkdir -p /opt/db_data
fi

chown postgres -R /opt/db_data

sudo -u postgres /usr/lib/postgresql/14/bin/initdb -D /opt/db_data

if [ $? -eq 0 ] 
then 
  echo "DB data has been inited."
  echo "host	all		all		0.0.0.0/0	trust" >> /opt/db_data/pg_hba.conf
  echo "listen_addresses = '*'" >> /opt/db_data/postgresql.conf
else 
  echo "DB data already exists and will be used without any modification."
fi

sudo -u postgres /usr/lib/postgresql/14/bin/postgres -D /opt/db_data -c config_file=/opt/db_data/postgresql.conf
