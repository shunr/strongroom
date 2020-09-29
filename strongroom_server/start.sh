#!/bin/bash

set -e

# Create DB and initialize tables
while true
do
    echo "SELECT 'CREATE DATABASE $POSTGRES_DB' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$POSTGRES_DB')\gexec" | psql \
        -U $POSTGRES_USER \
        -h db \
        --set ON_ERROR_STOP=on \
        --set DBNAME=\'$POSTGRES_DB\' && break
    sleep 5
done

PGPASSWORD=$POSTGRES_PASSWORD psql \
    -X \
    -U $POSTGRES_USER \
    -h db \
    -f init.sql \
    -d $POSTGRES_DB \
    --set ON_ERROR_STOP=on

go run github.com/shunr/strongroom_server --port $SERVER_PORT
