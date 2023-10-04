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

Running migrations via Docker:

```
cd docker && docker build -f db.Dockerfile -t ugentlib/projects-atlas ../
docker run -v $(pwd)/atlas.hcl:/atlas.hcl ugentlib/projects-atlas migrate apply --env local
```

## Development

### Live reload

```
go install github.com/cespare/reflex@latest
cp reflex.example.conf reflex.conf
reflex -c reflex.conf
```

## Database migrations

If you make a change to the schema files in `ent/schema/`, you will need to run these steps:

Re-generate the `ent` code with `cd ent && go generate ./...`. Next, generate a new migration
with `atlas migrate diff --env local`. This will generate a new migration file in 
`ent/migrate/migrations` and update the `atlas.sum` file in that directory. Finally, run the 
migration against your database with `atlas migrate apply --env local`.

### Data migrations

Create a new migration file with `atlas migrate new <name> --env local`. Then edit that file 
with `atlas migrate edit <filename> --env local`. This wil automatically update the `atlas.sum`
file. You can also edit the file directly, but then you have to run `atlas migrate hash --env local`
to re-generate the `atlas.sum` file.

## OpenAPI

The REST API is described through [OpenAPI](https://swagger.io/specification/). The code is generated
with [ogen](https://ogen.dev/).

When making changes to the API specification in `api/v1/openapi.yaml`, you must regenerate the API 
server code: `cd api/v1 && go generate ./...`.

The implementation of the handler resides in `api/v1/service.go`. Make sure all methods of the `Handler`
interface are implemented.
