#!/bin/sh

export GIN_MODE=debug
export BOLETIA_APP_MODE="develop"

# file logs
# export BOLETIA_PATH_FILE_LOGS="../gitignore/logs/all.log"

# http and https
export BOLETIA_ADDRESS_HTTP=":8080"
export BOLETIA_ADDRESS_HTTPS=":4443"
export BOLETIA_PATH_CERT_HTTPS="../gitignore/certs/cert.pem"
export BOLETIA_PATH_KEY_HTTPS="../gitignore/certs/key.pem"
export BOLETIA_READ_TIMEOUT="17"
export BOLETIA_WRITE_TIMEOUT="17"
# export BOLETIA_AUTHORIZATION_API_KEY="Asd*123." # api key authorization endpoints

# db
export BOLETIA_POSTGRES_DB_USER="boletia"
export BOLETIA_POSTGRES_DB_PASSWORD="boletia"
export BOLETIA_POSTGRES_DB_ADDRESS="localhost:5432" # servidor con su puerto
export BOLETIA_POSTGRES_DB_DBNAME="boletia"

export BOLETIA_DELAY_BETWEEN_CICLES="1" # cantidad de minutos q se demora en hacer otra llamada al la api

# currencyapi client
export BOLETIA_CURRENCY_API_URL_BASE="https://api.currencyapi.com/v3"
export BOLETIA_CURRENCY_API_KEY="rKNeGzWutRov1gT8KdrSIpoWrra7EOG0njtzoCav" 
export BOLETIA_CURRENCY_REQUEST_TIMEOUT="17" # timeout expresado en segundos

