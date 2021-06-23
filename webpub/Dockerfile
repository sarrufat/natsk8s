FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . ./
ARG GIT_VERSION
RUN go get github.com/sarrufat/natsk8s/pub/rand
RUN go get github.com/sarrufat/natsk8s/webpub/service
RUN go build -o /webpub -ldflags "-X main.version=$GIT_VERSION"


CMD [ "webpub" ]
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


COPY --from=build /webpub /webpub


USER nonroot:nonroot

ENTRYPOINT ["/webpub"]
EXPOSE 9090
