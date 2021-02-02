FROM golang:1.14-alpine as development

ENV CGO_ENABLED=1

WORKDIR /go/src/app

RUN apk update && apk add git
RUN go get github.com/cespare/reflex
COPY . .
RUN go build -o postgres-tool main.go

CMD ["reflex", "-c", "./reflex.conf"]

FROM alpine AS build
COPY --from=development /go/src/app/postgres-tool /opt/postgres-tool
ENTRYPOINT ["/opt/postgres-tool"]