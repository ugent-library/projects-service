# Project service

A service that publishes a directory of research projects at Ghent University.

## Setup

### Database

Create a new database and a user.

This application uses PostgreSQL's text search feature. You need to create 
a custom `TEXT CONFIGURATION` using these queries:

```sql
CREATE EXTENSION IF NOT EXISTS unaccent;
CREATE TEXT SEARCH CONFIGURATION usimple ( COPY = simple );
ALTER TEXT SEARCH CONFIGURATION usimple ALTER MAPPING FOR hword, hword_part, word WITH unaccent, simple;
```

The database is managed with [atlas](https://atlasgo.io/).

Copy the `atlas.hcl.example` to `atlas.hcl`. 

Atlas requires you to have two databases when developing, both are referenced in the `atlas.hcl` file:

* a [dev-database](https://atlasgo.io/concepts/dev-database) used to check the state 
of the database against the schema and migrations in order to generate a migration path.
* The target database containing data, on which migrations are applied.

In production, using a `dev-database` to verify the migration path and allow rollbacks is optional.

Initialize the database with the `env` flag:

```
atlas migrate apply --env local
```

### Environment variables

Copy `.env.example` to `.env` and ensure these variables are present:

```
PROJECTS_ENV               # environment (local, production, development, default:production)
PROJECTS_HOST              # host or IP (default: localhost)
PROJECTS_PORT              # host port (default: 3000)
PROJECTS_API_KEY           # REST API Key
PROJECTS_REPO_CONN         # PostgreSQL DSN connection string
PROJECTS_REPO_SECRET       # PostgreSQL secret seed
```

### Application boot

Manually:

```go
go run main.go server
```

Or via Docker:

```
cd docker && docker build -f app.Dockerfile -t ugentlib/projects ../
docker run ugentlib/projects /dist/app server
```

Use [tern](https://github.com/jackc/tern) to initalize the database:

```
cd migrations && tern migrate apply
```

Either create a `tern.conf` file in the `migrations` directory, or use `tern` with 
[PG environment variables](https://www.postgresql.org/docs/current/libpq-envars.html).

## Development

### Live reload

```
go install github.com/cespare/reflex@latest
cp reflex.example.conf reflex.conf
reflex -c reflex.conf
```

## Database migrations

If you make a change to the schema files in `ent/schema/`, you will need to run these steps:

* `cd ent && go generate ./...` to generate `ent` Go code.
* `tern migrate new` to create a new migration file.
* Add SQL code to the migration file.
* `tern migrate apply` to apply pending migrations to the database.

**Experimental: atlas**

[atlas](https://atlasgo.io) is a companion to `ent`. It calculates a migration path by diffing
the schema in Go against the migrations. It then generates a new migration in SQL. This project
supports atlas, but prefers tern as atlas relies on [lib/pq](https://github.com/lib/pq).

## OpenAPI

The REST API is described through [OpenAPI](https://swagger.io/specification/). The code is generated
with [ogen](https://ogen.dev/).

When making changes to the API specification in `api/v1/openapi.yaml`, you must regenerate the API 
server code: `cd api/v1 && go generate ./...`.

The implementation of the handler resides in `api/v1/service.go`. Make sure all methods of the `Handler`
interface are implemented.
