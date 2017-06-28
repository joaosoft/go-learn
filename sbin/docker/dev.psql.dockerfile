FROM postgres:9.5

ADD ./schema/bootstrap/* /docker-entrypoint-initdb.d/
ADD ./services/mirakl-mci/schema/01.sql  /docker-entrypoint-initdb.d/mirakl-mci-01.sql
ADD ./services/mirakl/schema/01.sql  /docker-entrypoint-initdb.d/mirakl-01.sql
ADD ./services/products/schema/01.sql /docker-entrypoint-initdb.d/products-01.sql
ADD ./services/retek/schema/01-integration.sql  /docker-entrypoint-initdb.d/retek-01-integration.sql
ADD ./services/sellers/schema/01.sql  /docker-entrypoint-initdb.d/sellers-01.sql
