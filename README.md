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

### Environment variables

Copy `.env.example` to `.env` and ensure these variables are present:

```
PROJECTS_ENV               # environment (local, production, development, default:production)
PROJECTS_HOST              # host or IP (default: localhost)
PROJECTS_PORT              # host port (default: 3000)
PROJECTS_API_KEY           # REST API Key
PROJECTS_REPO_CONN         # PostgreSQL DSN connection string
PROJECTS_SEARCH_CONN       # Search (es6) DSN connection string
PROJECTS_SEARCH_INDEX      # Search (es§) index name
```

### Application boot

Via Reflex:

```go
cp reflex.conf.example .reflex.conf
reflex -d none -c .reflex.conf
```

Or via Docker:

```
cd docker && docker build -f app.Dockerfile -t ugentlib/projects ../
docker run ugentlib/projects /dist/app server
```

Use [tern](https://github.com/jackc/tern) to initalize the database:

```
cd etc/migrations && tern migrate apply
```

Either create a `tern.conf` file in the `etc/migrations` directory, or use `tern` with 
[PG environment variables](https://www.postgresql.org/docs/current/libpq-envars.html).

## Development

### Live reload

```
go install github.com/cespare/reflex@latest
cp reflex.example.conf reflex.conf
reflex -c reflex.conf
```

## Database

Making changes to the database schema:

* `tern migrate new` to create a new migration file.
* Add SQL code to the migration file.
* `tern migrate apply` to apply pending migrations to the database.

This projects uses [sqlc](https://sqlc.dev/).

## OpenAPI

The REST API is described through [OpenAPI](https://swagger.io/specification/). The code is generated
with [ogen](https://ogen.dev/).

When making changes to the API specification in `api/v1/openapi.yaml`, you must regenerate the API 
server code: `cd api/v1 && go generate ./...`.

The implementation of the handler resides in `api/v1/service.go`. Make sure all methods of the `Handler`
interface are implemented.