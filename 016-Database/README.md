# More about DataBase

## Notes

- On Dockerfile there is every installation step.

## Migrations

We can use a lib called golang-migrate. [Doc](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Commands

- Create a new migration: `migrate create -ext sql -dir sql/migrations -seq name_of_migration`
- Run migrations: `migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up`
- Rollback migrations: `migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down`
