## Postgres Setup

[This article](https://aws.amazon.com/blogs/database/managing-postgresql-users-and-roles) outlines fine-grained access control in PostgreSQL.

### Tasks
- Create database: `CREATE DATABASE tenantwin;`.
- Create tables for each entity in `documentation/diagrams/entities` diagram: `./tables.sql`.
- Create users for each service in `documentation/diagrams/services` diagram. Create a `<service_name>_readwrite` role for each service with the required permissions on table(s), schema and database. Grant each service its readwrite role. All steps in: `./users-roles.sql` (modify passwords first).