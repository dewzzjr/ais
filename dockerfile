FROM alpine:latest

WORKDIR /srv
ADD bin/server ./server

EXPOSE 9000

COPY --from=ghcr.io/ufoscout/docker-compose-wait:latest /wait /wait

CMD /wait && ./server