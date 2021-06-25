module github.com/sarrufat/natsk8s/webpub

go 1.16

replace github.com/sarrufat/natsk8s/webpub/service => ./service

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/sarrufat/natsk8s/webpub/service v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.3.0
)
