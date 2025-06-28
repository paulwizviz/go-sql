# Data Model

This section contains a list of data models (logical amd schema) we'll use a basis to implement as working examples.

## Example 1: Person Profile Domain

This example describes the relationship between a person and home address.

### Example 1 data schema

* Person
  * Primary key: id
  * FirstName: Text
  * MiddleName: Text
  * Surname: Text

* Person Profile
  * Primary key: id
  * Foreign key: person_id
  * Email: Text

* Location
  * Primary key: id
  * Foreign key: person_id
  * Address: Text

## Example 2: Property Ownership Domain

This example is based on property ownership domain.

### Relational schema
