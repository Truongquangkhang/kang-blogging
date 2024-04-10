# sql

Store the DB schema and DB migration scripts

### Directories

Each Service will have a package to store DB schema and DB migrations

```
/service-1
  /migration
    000000_do_something.down.sql
    000000_do_something.up.sql
    ...
  schema.dbml
...
```

Current Directories:

- [nio-core](nio-core/) DB schema and DB migration for NIO Core service

### To create a migration

We are using [Go Migrate](https://github.com/golang-migrate/migrate) for DB migration. To create a new migration, install as a go module and run:

> migrate create -ext sql -dir sql/[service_name]/migration -seq [migration_name]

**Note**:
- We will create a new file with `sql` extension: `-ext sql`
- We will use **6 digits number** instead of timestamp: `-seq`
- Must put the migration in the service migration package: `-dir sql/[service_name]/migration`

And make sure to update the `schema.dbml` to the latest DB schema
