FROM ubuntu:18.04 AS runtime
WORKDIR /ogwc
COPY build/ogwc /bin/ogwc
EXPOSE 16666
ENV REDIS_ADDR=myredis:6379
ENV REDIS_DB=0
ENV REDIS_PW=""
ENV BIND_ADDR=":16666"
CMD ogwc server --bind-address="$BIND_ADDR" --redis-address="$REDIS_ADDR" --redis-database="$REDIS_DB" --redis-password="$REDIS_PASSWORD"