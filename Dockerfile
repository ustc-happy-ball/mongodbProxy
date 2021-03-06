FROM golang:1.15-alpine as builder
WORKDIR /root/go/src/github.com/LILILIhuahuahua/ustc_tencent_game
COPY . /root/go/src/github.com/LILILIhuahuahua/ustc_tencent_game
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
#RUN apk add --no-cache --virtual .build-deps gcc musl-dev
#RUN git config --global url."".insteadOf ""
#RUN export GOPRIVATE=git.enjoymusic.ltd && go build -o bifrost-api main.go plugin.go
RUN go build -o db-svc main.go

FROM alpine:latest
# environment variable for mongoDB connection
ARG DB_USER
ARG DB_PWD
ARG DB_HOST
ARG DB_PORT
ENV  ENV_DB_USER=$DB_USER \
     ENV_DB_PWD=$DB_PWD \
     ENV_HOST=$DB_HOST \
     ENV_PORT=$DB_PORT

WORKDIR  /root/go/src/github.com/LILILIhuahuahua/ustc_tencent_game
COPY --from=builder  /root/go/src/github.com/LILILIhuahuahua/ustc_tencent_game/db-svc .
EXPOSE 8890/tcp
ENTRYPOINT ./db-svc -DBUser $ENV_DB_USER -DBPassword $ENV_DB_PWD -Host $ENV_HOST -Port $ENV_PORT
#ENTRYPOINT ["./db-svc", "-DBUser", "${ENV_DB_USER}","-DBPassword", "${ENV_DB_PWD}","-Host","${ENV_HOST}","-Port", "${ENV_PORT}"]