clean:
	go clean

build:
	.tools/build.sh default

build.run:
	.tools/build.sh default && ./.build/hepburn web -c config.hcl

build-linux:
	.tools/build.sh linux

test:
	go test ./...

lint:
	golangci-lint -v run

depends-up:
	docker-compose up -d --remove-orphans

depends-stop:
	docker-compose stop
