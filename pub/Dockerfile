FROM golang:1.16-alpine AS build

WORKDIR /app

COPY *.go ./
COPY go.* ./

RUN go build -o /nats-pub


CMD [ "nats-pub" ]
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


COPY --from=build /nats-pub /nats-pub


USER nonroot:nonroot

ENTRYPOINT ["/nats-pub"]