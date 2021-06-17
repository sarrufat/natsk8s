FROM golang:1.16-alpine AS build

WORKDIR /app

COPY *.go ./
COPY go.* ./

RUN go build -o /nats-sub


CMD [ "nats-sub" ]
##
## Deploy
##

FROM alpine

ARG USER=nonroot
ENV HOME /home/$USER

# install sudo as root
RUN apk add --update sudo

# add new user
RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME


COPY --from=build /nats-sub /nats-sub


USER nonroot:nonroot

ENTRYPOINT ["/nats-sub"]