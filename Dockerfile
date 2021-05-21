 FROM registry.semaphoreci.com/golang:1.16 as builder
 ENV APP_USER app
 ENV APP_HOME /go/src/intranet
 RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
 RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
 WORKDIR $APP_HOME
 USER $APP_USER
 COPY src/ .
 RUN go mod download
 RUN go mod verify
 RUN go build -o intranet
 FROM debian:buster
 FROM registry.semaphoreci.com/golang:1.16
 ENV APP_USER app
 ENV APP_HOME /go/src/intranet
 RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
 RUN mkdir -p $APP_HOME
 WORKDIR $APP_HOME
 COPY src/site/ site/
 COPY --chown=0:0 --from=builder $APP_HOME/intranet $APP_HOME
 EXPOSE 8089
 USER $APP_USER
 CMD ["./intranet"]