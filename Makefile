BINARY=endpoint
BINARY_CONSUMER=visitor

PACKAGE=github.com/nickylogan/guestbook

vendor:
	@dep ensure -v

build:
	GOOS=linux go build -o ./cmd/${BINARY}/${BINARY} ${PACKAGE}/cmd/${BINARY}

build_consumer:
	GOOS=linux go build -o ./cmd/${BINARY_CONSUMER}/${BINARY_CONSUMER} ${PACKAGE}/cmd/${BINARY_CONSUMER}

clean:
	if [ -f ./cmd/${BINARY}/${BINARY} ] ; then rm ./cmd/${BINARY}/${BINARY} ; fi
	if [ -f ./cmd/${BINARY_CONSUMER}/${BINARY_CONSUMER} ] ; then rm ./cmd/${BINARY_CONSUMER}/${BINARY_CONSUMER} ; fi