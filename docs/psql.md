# Postgres

PostgreSQL is an object-relational database management system (ORDBMS) based on POSTGRES, Version 4.2, developed at the University of California at Berkeley Computer Science Department. POSTGRES pioneered many concepts that only became available in some commercial database systems much later[1](https://www.postgresql.org/docs/current/intro-whatis.html).

## Postgres client application

Here is the [list](https://www.postgresql.org/docs/current/reference-client.html)

## PgAdmin

* [Container setup](https://www.pgadmin.org/docs/pgadmin4/latest/container_deployment.html)

## The SQL Language

Here is the [official documentation](https://www.postgresql.org/docs/current/sql.html)

## A test network

A docker compose based network is found [here](../deployment/postgres/docker-compose.yml). In this setup you will find a postgres server (hostname: server).

## Operational scripts

A bash shell script named [./scripts/postgres.sh](../scripts/postgres.sh) is provided to help you work operate the network. The operations included in the scripts are:

* start server - `./scripts/postgres.sh start`
* stop server  - `./scripts/postgres.sh stop`
* clean setup  - `./scripts/postgres.sh clean`
* build client - `./scripts/postgres.sh build`
* start client - `./scripts/postgres.sh cli`

## Setting up the network

The steps for setting up the test network are:

* STEP 1 - Clean setup
* STEP 2 - Build cli image
* STEP 3 - Start server

## Using postgres client

Access the client by 

When you in the postgres cli, you have two ways of accessing the server.

Option 1:

```sh
psql -h [HOSTNAME] -p [PORT] -U [USERNAME] -W -d [DATABASENAME]
```

Option 2:


```sh
psql postgres://[USERNAME]:[PASSWORD]@[HOSTNAME]:[PORT]/[DATABASENAME]?sslmode=require
```

In the context of the working example, the environment variables are listed as environment variables:

* USERNAME - postgres
* PASSWORD - postgres
* HOSTNAME - server (docker compose service name)
* DATABASENAME - psql_db