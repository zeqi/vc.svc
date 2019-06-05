
BINARY = vc-svc
VERSION=0.1.0
GITHUB_USERNAME=zeqi
GITHUB_REPO=${GITHUB_USERNAME}/${BINARY}
DOCKER_REPO=zeqi/$(BINARY)
IMAGE_NAME=${DOCKER_REPO}:${VERSION}

LDFLAGS = -ldflags "-X main.VERSION=${VERSION}"
pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)
gobuild = go build ${LDFLAGS} -o ${BINARY}

plugin = plugin.go
server = main.go


all: format build-linux-server

format:
	go fmt $(pkgs)

test:
	go test `go list ./... | grep -v apis | grep -v client`

# gen-proto:
# 	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

build:
	$(gobuild) $(server) $(plugin)

build-linux-server:
	CGO_ENABLED=0 GOOS=linux $(gobuild) $(server) $(plugin)

docker-build:
	docker rmi $(IMAGE_NAME)
	# docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
	docker build . -t $(IMAGE_NAME)

# For testing purposes
# consul-start:
# 	mkdir /tmp/consul || true
# 	consul agent -server -bootstrap-expect=1 -data-dir=/tmp/consul -advertise=127.0.0.1 > /tmp/consul/consul.log 2>&1 &