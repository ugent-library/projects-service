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

## Dev Containers

This project supports [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers). Following these steps
will auto setup a containerized development environment for this project. In VS Code, you will be able to start a terminal
that logs into a Docker container. This will allow you to write and interact with the code inside a self-contained sandbox.

**Installing the Dev Containers extension**

1. Open VS Code.
2. Go to the [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension page.
3. Click the `install` button to install the extension in VS Code.

**Open in Dev Containers**

1. Open the project directory in VS Code.
2. Click on the "Open a remote window" button in the lower left window corner.
3. Choose "reopen in container" from the popup menu.
4. The green button should now read "Dev Container: App name" when successfully opened.
5. Open a new terminal in VS Code from the `Terminal` menu link.

You are now logged into the dev container and ready to develop code, write code, push to git or execute commands.

**Run the project**

1. Open a new terminal in VS Code from the `Terminal` menu link.
2. Execute this command `reflex -d none -c reflex.docker.conf`.
3. Once the application has started, VS Code will show a popup with a link that opens the project in your browser.

**Networking**

The application and its dependencies run on these ports:

| Application      | Port |
| ---------------- | ---- |
| Projects Service | 3301 |
| DB Application   | 3351 |
| Elastic Search   | 3361 |