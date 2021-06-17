FROM golang:1.16-alpine

WORKDIR /app

COPY *.go ./
COPY go.* ./

RUN go build -o /nats-sub


CMD [ "nats-sub" ]
##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /nats-sub /nats-sub


USER nonroot:nonroot

ENTRYPOINT ["/nats-sub"]