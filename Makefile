BINARY_NAME=otc-app
SECRETKEY=""

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BINARY_NAME} .
	docker build -t ${BINARY_NAME} .

build-docker-compose:
	docker-compose build

run-docker-compose:
	SECRET_KEY=$(SECRETKEY) docker-compose up -d

stop-docker-compose:
	docker-compose down

run:
	./${BINARY_NAME}

build_and_run: build run

test:
	go test -v

clean:
	go clean
	rm -f ${BINARY_NAME}
