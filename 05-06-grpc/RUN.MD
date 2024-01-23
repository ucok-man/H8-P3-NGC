# FTGO-P3-V1-LC1

## ENVIRONMENT

Required:

- GATEWAY_PORT
- USER_SERVICE_PORT
- USER_SERVICE_DB_DSN
  - atlas: mongodb+srv://user:password@url/dbname?retryWrites=true&w=majority&authSource=admin
  - local: mongodb://user:password@host:port/dbname?retryWrites=true&w=majority&authSource=admin
- USER_SERVICE_DB_NAME
- JWT_SECRET

Optional:

- USER_SERVICE_DB_MIN_POOL_SIZE
- USER_SERVICE_DB_MAX_POOL_SIZE
- USER_SERVICE_DB_MAX_IDLE_TIME (ex: 15m)

## HOW TO RUN

<!-- Manual:

- Make go/start

Docker:

- Make docker/container/run name=(container name)
- docker container run --name (containername) --publish (hostport):(containerport) --detach -e API_PORT="" -e DB_DSN="" -e DB_NAME="" jumatberkah/livecode-1-p3 -->

<!-- ### Docker Registry -->

<!-- [https://hub.docker.com/r/jumatberkah/livecode-1-p3](https://hub.docker.com/r/jumatberkah/livecode-1-p3) -->
