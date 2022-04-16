# Doc Enviroment

## configuration
- GIN_MODE ==> Modo despliegue de gin framework. Valores: debug, release. Default: debug.
- COSMOZ_APP_MODE ==> Modo de despliegue de la app. Valores: develop, production. Default: develop.

## logs
- COSMOZ_PATH_FILE_LOGS ==> Ruta de un archivo para guardar los logs. Si esta variable NO esta presente en el enviroment o su valor es vacio los logs se muestran por consola. Default:.

## http y https
- COSMOZ_ADDRESS_HTTP ==> Host y puerto para http. Default: :8080.
- COSMOZ_ADDRESS_HTTPS ==> Host y puerto para https. Default::4443.
- COSMOZ_PATH_CERT_HTTPS ==> Ruta certificado ssl para https. Si esta variable NO esta presente en el enviroment NO se levanta el servidor HTTPS. Default: .
- COSMOZ_PATH_KEY_HTTPS ==> Ruta de la clave ssl para https. Si esta variable NO esta presente en el enviroment NO se levanta el servidor HTTPS. Default: .
- COSMOZ_READ_TIMEOUT ==> Tiempo de lectura en segundos para una peticion HTTP o HTTPS. Default: 10.
export COSMOZ_WRITE_TIMEOUT ==> Tiempo de escritura en segundos para una peticion HTTP o HTTPS. Default: 10.

## bd
- COSMOZ_DB_MONGO_USER ==> Usuario para acceso a la base de datos.
- COSMOZ_DB_MONGO_PASSWORD ==> Clave para acceso a la base de datos.
- COSMOZ_DB_MONGO_ADDRESS  ==> Servidor de base de datos.
- COSMOZ_DB_MONGO_DBNAME  ==> Nombre de la base de datos.

## autorizacion
- COSMOZ_ACCESS_TOKEN_SECRET_KEY ==> Clave para generar el JWT.
- COSMOZ_ACCESS_TOKEN_AUDIENCE ==> Audiencia del JWT
- COSMOZ_ACCESS_TOKEN_ISSUER ==> Emisor del JWT
- COSMOZ_ACCESS_TOKEN_LIVE ==> Duracion en segundos q el JWT es valido.
- COSMOZ_REFRESH_TOKEN_SECRET_KEY ==> Clave para generar el refresh-token.
- COSMOZ_REFRESH_TOKEN_LIVE ==> Duracion en segundos q el refresh-token es valido.

## upload files(medias,...)
- COSMOZ_URL_SERVER_STORE_MEDIAS ==> Url para acceder a los archivos q se subieron(avatar, ...)
- COSMOZ_STORE_UPLOADED_FILES_MODE ==> Donde se suben los archivos. Values:  files-system-local, aws-s3. Default: files-system-local.
- COSMOZ_DESTINATION_TARGET ==> Ruta o Bucket S3 AWS donde se suben los archivos. Ruta si COSMOZ_STORE_UPLOADED_FILES_MODE=files-system-local, Nombre del Bucket S3 si COSMOZ_STORE_UPLOADED_FILES_MODE=aws-s3.