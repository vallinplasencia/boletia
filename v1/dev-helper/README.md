# Doc Enviroment

## configuration
- GIN_MODE ==> Modo despliegue de gin framework. Valores: debug, release. Default: debug.
- BOLETIA_APP_MODE ==> Modo de despliegue de la app. Valores: develop, production. Default: develop.

## logs
- BOLETIA_PATH_FILE_LOGS ==> Ruta de un archivo para guardar los logs. Si esta variable NO esta presente en el enviroment o su valor es vacio los logs se muestran por consola. Default:.

## http y https
- BOLETIA_ADDRESS_HTTP ==> Host y puerto para http. Default: :8080.
- BOLETIA_ADDRESS_HTTPS ==> Host y puerto para https. Default::4443.
- BOLETIA_PATH_CERT_HTTPS ==> Ruta certificado ssl para https. Si esta variable NO esta presente en el enviroment NO se levanta el servidor HTTPS. Default: .
- BOLETIA_PATH_KEY_HTTPS ==> Ruta de la clave ssl para https. Si esta variable NO esta presente en el enviroment NO se levanta el servidor HTTPS. Default: .
- BOLETIA_READ_TIMEOUT ==> Tiempo de lectura en segundos para una peticion HTTP o HTTPS. Default: 10.
- export BOLETIA_WRITE_TIMEOUT ==> Tiempo de escritura en segundos para una peticion HTTP o HTTPS. Default: 10.

## bd
- BOLETIA_POSTGRES_DB_USER ==> Usuario para acceso a la base de datos.
- BOLETIA_POSTGRES_DB_PASSWORD ==> Clave para acceso a la base de datos.
- BOLETIA_POSTGRES_DB_ADDRESS  ==> Servidor de base de datos.
- BOLETIA_POSTGRES_DB_DBNAME  ==> Nombre de la base de datos.

## currencyapi client
- BOLETIA_CURRENCY_API_URL_BASE ==> Url base de currencyapi.
- BOLETIA_CURRENCY_API_KEY ==> Api key para acceder a currencyapi
- BOLETIA_CURRENCY_REQUEST_TIMEOUT ==> Tiempo maximo(segunods) q se puede estar haciendo la peticion a la currecyapi

## ciclos
- BOLETIA_DELAY_BETWEEN_CICLES ==> Tiempo(minutos) q se demora entre peticion y peticion a currencyapi