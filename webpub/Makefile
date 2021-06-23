

GIT_COMMIT=$(shell git rev-list -1 HEAD)


build-docker:
	docker build --build-arg GIT_VERSION=$(GIT_COMMIT) -t localhost:32000/webpub .
	docker push localhost:32000/webpub
