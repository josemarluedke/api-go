# API

Api written in Golang.


## Usage

Start the server:

```gin run server.go```

## Migrations

We use [go-pg/migrations](https://github.com/go-pg/migrations) to control our database migrations, such as creating, changing and dropping tables & fields.

There will be a table called `gopg_migrations` that saves all the executed migrations.

To setup the migrations table, execute:

```go run db/migrations/*.go init```

To run migrations, execute:

```go run db/migrations/*.go up```

To revert the last migration, execute:

```go run db/migrations/*.go ```
