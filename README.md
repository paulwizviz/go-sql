# Overview

Welcome to my collection of educational materials created by me and others about SQL Database. Included in this project are working examples, mostly writing in Go, and using Docker and Kubernates for executation. If you wish to work with these examples, please install:

* Go version 1.18 or later
* Docker including compose
* MiniKube or provide your own version of Kubernetes

## Topics

* [SQLite](./docs/sqlite.md)
* [Postgres](./docs/psql.md)

## References

* [Managing connections](https://go.dev/doc/database/manage-connections)
* [SQLite vs MySQL vs PostgreSQL: A Comparison Of Relational Database Management Systems](https://www.digitalocean.com/community/tutorials/sqlite-vs-mysql-vs-postgresql-a-comparison-of-relational-database-management-systems)

## Project structure

* `build` -- scripts used to create apps and containers.
* `cmd` -- Go code to build apps
* `deployment` -- docker and/or k8s network
* `doc` -- markdowns to complement README.md
* `scripts` -- mainly bash scripts to trigger build and deployment operations.

## Disclaimers

The working examples in this projects are purely for illustration only and are subject to modification without notice. Any opinions expressed is this project mine or belongs to the author of any referenced materials.

## Copyright

Unless otherwise specificed, copyright of the working examples are assigned as follows.

Copyright 2022 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.