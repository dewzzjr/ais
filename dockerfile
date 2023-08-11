FROM alpine:3.18

WORKDIR /srv
ADD bin/server ./server

EXPOSE 9000

COPY --from=ghcr.io/ufoscout/docker-compose-wait:2.12.0 /wait /wait

CMD /wait && ./server