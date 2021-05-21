FROM golang:1.16.0-alpine as build

WORKDIR /go/src
COPY / /go/src

RUN go get
RUN go install
RUN go build

FROM alpine
COPY --from=build /go/src/dashboard /
COPY --from=build /go/src/site /site
CMD ["/dashboard"]
