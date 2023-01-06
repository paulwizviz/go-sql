# PostgreSQL

PostgreSQL is an object-relational database management system (ORDBMS) based on POSTGRES, Version 4.2, developed at the University of California at Berkeley Computer Science Department. POSTGRES pioneered many concepts that only became available in some commercial database systems much later[1](https://www.postgresql.org/docs/current/intro-whatis.html).

## PostgreSQL Applications

<u>Client utility applications</u>

The common feature of these applications is that they can be run on any host, independent of where the database server resides.

Please refer to this [list](https://www.postgresql.org/docs/current/reference-client.html)

<u>Server applications</u>

These commands can only be run usefully on the host where the database server resides.

Please refer to this [list](https://www.postgresql.org/docs/current/reference-server.html)

<u>pgAdmin</u>

An open source adminstration platform for PostgreSQL

* [Container setup](https://www.pgadmin.org/docs/pgadmin4/latest/container_deployment.html)

## SQL Command

Here is the [official documentation](https://www.postgresql.org/docs/current/sql.html)

## Working Examples 

The setup for working examples are:

* [Operational script](../scripts/postgres.sh)
* [Docker compose deployment](../deployment/postgres/docker-compose.yml).
* Shell container

<u>Operational script</U>

This is a bash script to help you manage the working example.

<u>Docker compose deployment</u>

The steps to operate the test network are:

* Start network - `./scripts/psql.sh network start`
* Stop network - `./scripts/psql.sh network stop`

<u>Shell container</u>

The working examples provide an Ubuntu based container with psql client applications installed. To use it, please follow these steps:

* STEP 1 - If you had not already done so, build this image. Run this command `./scripts/psql.sh image build`.
* STEP 2 - Access the shell container, run this command `./scripts/psql.sh client cli`


### Example 1 - Connecting to server via psql

 Using the `psql` app to connect to the default server.

* USERNAME: postgres
* PASSWORD: postgres
* HOSTNAME: defaultserver // docker compose container name
* PORT: 5432
* DATABASENAME: default

Option 1:

```sh
psql -h [HOSTNAME] -p [PORT] -U [USERNAME] -W -d [DATABASENAME]
```

Option 2:

```sh
psql postgres://[USERNAME]:[PASSWORD]@[HOSTNAME]:[PORT]/[DATABASENAME]?sslmode=require
```

### Example 2 - Using pgAdmin for DevOp

* email: admin@psql.email
* password: admin
* port: 5050

* STEP 1: Start network `./scripts/psql.sh network start`
* STEP 2: Open browser `http://localhost:5050/browser/`

### Example 3 - Using client app for DevOp

Please refer to this [script](../db/psql/scripts/testdb.sh)

* STEP 1: Start network `./scripts/psql.sh network start`
* STEP 2: Open cli `./scripts/psql.sh client cli`
* STEP 3: In the cli, create db `./scripts/testdb.sh db`
* STEP 4: In the cli, create schema `./scripts/testdb.sh schema` 