BINARY=endpoint

PACKAGE=github.com/nickylogan/guestbook

vendor:
	@dep ensure -v

build:
	GOOS=linux go build -o ./cmd/${BINARY}/${BINARY} ${PACKAGE}/cmd/${BINARY}

clean:
	if [ -f ./cmd/${BINARY}/${BINARY} ] ; then rm ./cmd/${BINARY}/${BINARY} ; fi

run:
	sudo docker-compose -f "./deployments/docker/docker-compose.yaml" up
