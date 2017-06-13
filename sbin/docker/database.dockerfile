FROM postgres:9.5

ADD ./schemas/create.sql /docker-entrypoint-initdb.d/
