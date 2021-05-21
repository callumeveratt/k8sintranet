FROM golang:1.16.0-alpine as build

WORKDIR /go/src
COPY /src /go/src

RUN go get
RUN go install
RUN go build

FROM alpine
COPY --from=build /go/src/intranet /
CMD ["/intranet"]