# MySQL

Please refer to [MySQL Home](https://dev.mysql.com/doc/)

## MySQL Applications

* [phpMyAdminer](https://www.adminer.org/)

* The MySQL Command-Line Client
    * [v8](https://dev.mysql.com/doc/refman/8.0/en/mysql.html)

## Working Examples

Resources:

* [Operational script](../scripts/mysql.sh)
* [Docker compose deployment](../deployment/mysql/docker-compose.yml).
* Shell container

<u>Operational script</U>

This is a bash script to help you manage the working example.

<u>Docker compose deployment</u>

The steps to operate the test network are:

* Start network - `./scripts/mysql.sh network start`
* Stop network - `./scripts/mysql.sh network stop`

<u>Shell container</u>

The working examples provide an Ubuntu based container with mysql client applications installed. To use it, please follow these steps:

* STEP 1 - If you had not already done so, build this image. Run this command `./scripts/mysql.sh image build`.
* STEP 2 - Access the shell container, run this command `./scripts/mysql.sh client cli`

### Example 1 - Connecting to server via mysql

```sh
mysql -h defaultserver -u root -p
password:
```

### Example 2 - Connecting to server via adminer

**HOSTS**: defaultserver
**USERNAME**: root
**PASSWORD**: example

* STEP 1: Start network `./scripts/mysql.sh network start`
* STEP 2: Open browser `http://localhost:8080/`

