FROM arigaio/atlas:latest

VOLUME /atlas.hcl

COPY ent/migrate/migrations /ent/migrate/migrations