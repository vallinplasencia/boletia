# Boletia

Boletia app

## Tecnologia

- [Golang 1.16](https://golang.org)
- [Gin framework 1.7.4](https://github.com/gin-gonic/gin)
- [PostgresSQL 14.1](https://www.postgresql.org)

## Desplegar aplicacion

Clonar el repo:
~~~ sh
    # clonar el repo
    git clone https://github.com/vallinplasencia/boletia
~~~
Levantar la app se pude hacer de diferentes formas:
1. A traves de docker-compose:
~~~ sh
    # directorio de script
    cd v1/dev-helper
    # generar certificados para el endpoint web
    ./ssl-generator.sh
    # directorio de archivos de configuracion de docker-compose
    cd ../v1/docker-compose
    # copia script sql a un volumen para postgres
    ./pre-deploy.sh
    # levanta la app y los servicios q ella requiere.
    ./up.sh
~~~
2. Manualmente

    - Levantar servidor de bd postgreSQL.
    - Crear una nueva base de datos preferentemente con el nombre "**boletia**
    - Importar archivo .sql q se encuentra en **v1/docker-compose/misc/postgres/init.sql**.
    - Configurar las variables de acceso a la base de datos en el archivo **v1/dev-helper/env-local.sh**
    - Correr el script **./ssl-generator.sh** q se encuentra en **v1/dev-helper/env-local.sh** para generar los certificados ssl.
    - Correr la app entrando al directorio **v1/dev-helper/** y ejecutando **./run-server.sh**.

## Configuracion

EL punto de entrada de la app radica en **v1/cmd/main.go** y puede ejecutarla navegando al directorio **v1/dev-helper/** y corriendo **./run-server.sh** o a traves de **docker-compose** en el directorio **v1/docker-compose/**. Ver seccion: Desplegar aplicacion.


Toda la configuracion q necesita la app se realiza a traves de variables de entorno. Las cuales estan documentadas en el directorio **v1/dev-helper/README.md**

Los endpoints de la app pueden ser accedidos a traves de HTTP o HTTPS. Para generar certificado y clave para HTTPS navegar al directori **v1/dev-helper/** y correr **./ssl-generator.sh**.

## BD

- Script sql de la bd, se encuentra en **v1/docker-compose/misc/postgres/init.sql**.