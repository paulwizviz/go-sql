# SQL Schema

This section contains a list of data models (logical amd schema) we'll use a basis to implement as working examples.

## Example 1: Person Profile Domain

This example is a SQL implementation based on the domain model about a [Person Identity System](https://github.com/paulwizviz/learn-clean-architecture/blob/main/domains/person-profile.md).

NOTE: The focus of our implementation is the relationship between `person` and `name_identifier`

The table and fields are:

* Table `person`
  * Primary key: `id`

* Table `person_named_identifier`
  * Primary key: `id`
  * Foreign key: `person_id` Reference `person (id)`
  * Foreign key: `named_identifier_id` Reference `named_identifier (id)`

* Table `named_identifier`
  * Primary key: `id`
  * TEXT (not null): `first_name`
  * TEXT (not null): `surname`
  * TEXT (not null): `nickname`
