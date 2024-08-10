# More about DataBase

## Notes

- On Dockerfile there is every installation step.

## Migrations

We can use a lib called golang-migrate. [Doc](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Commands

- Create a new migration: `migrate create -ext sql -dir sql/migrations -seq name_of_migration`
- Run migrations: `migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up`
- Rollback migrations: `migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down`

## SQLX

SQLX is a lib that helps us to work with SQL in Go. [Doc](https://github.com/jmoiron/sqlx)

## SQLC

SQLC is a lib that helps us to generate Go code from SQL queries. [Doc](https://sqlc.dev/)

Support for MySQL, PostgreSQL, and SQLite.

### Usage

- Create a sqlc.yaml file with the specifications.
  - Schema: Path to the schema file.
  - Queries: Path to the queries file.
  - Engine: The database engine.
  - Gen: Path and language to generate the code.
- Run `sqlc generate` to generate the code.

Now we can use the generated code to interact with the database.
