## Postgres Setup

### Installation and Database Setup

All of the following steps are commands to run in terminal:

1. `brew install postgresql`
2. `brew services start postgresql`
3. `psql postgres`
4. `CREATE DATABASE tenantwin;`
5. `\c tenantwin;`
6. `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` // to allow for uuidv4 primary keys


### Tables, Users and Roles

[This article](https://aws.amazon.com/blogs/database/managing-postgresql-users-and-roles) outlines fine-grained access control in PostgreSQL.

More terminal commands, and assumes already connected to the `tenantwin` database.

1. Create tables for each entity in `documentation/diagrams/entities` diagram: `\i /<absolute-path>/tables.sql`
- Create users for each service in `documentation/diagrams/services` diagram. Create a `<service_name>_readwrite` role for each service with the required permissions on table(s), schema and database. Grant each service its readwrite role. All steps in: `\i /<absolute-path>/users-roles.sql` (modify passwords first).