# Project service

A service that publishes a directory of research projects at Ghent University.

## Development

### Live reload

```
go install github.com/cespare/reflex@latest
cp reflex.example.conf reflex.conf
reflex -c reflex.conf
```

### REST API

### Nats

First, install the [nats client](). Then, start a NATS server in development mode using the client.
Finally, create a stream and a consumer

```
nats server run --jetstream
nats stream add
nats consumer add
```

The output of `nats server` will list credentials for an auto-generated `user` account.
Use the credentials to configure the connection with the `PROJECTS_NATS_URL` env variable.

The consumer and stream should match these values in `.env`:

```
PROJECTS_NATS_STREAM
PROJECTS_NATS_CONSUMER
```

Start the importer:

```
go run main.go import
```

Publish messages to the stream like this:

```
nats pub project.a --count=3 --sleep 1s '{
    "type": "ResearchProject",
    "name": "Project: title",
    "foundingDate": "2015-01-01",
    "dissolutionDate": "2016-01-01",
    "description": "lorem ipsum dolor sit amet.",
    "identifier": [
        {
            "type": "PropertyValue",
            "propertyID": "IWETO",
            "value":  "00B99999"
        },
        {
            "type": "PropertyValue",
            "propertyID": "GISMO",
            "value":  "f0b38757-f375-4d24-b6c4-879955f8589a"
        }
    ]
}'
```
