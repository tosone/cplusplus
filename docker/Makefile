app_name                := cplusplus-build
docker_name             := $(app_name)
docker_tag              := dev
docker_container        := $(app_name)

.PHONY: upgrade
upgrade:
	docker pull golang:1.11.0-alpine3.8

.PHONY: build
build:
	docker build -t $(docker_name):$(docker_tag) .

.PHONY: run
run:
	docker-compose up --force-recreate --build -d

.PHONY: exec
exec:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish

.PHONY: exec-build
exec-build:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish -c \
	"cd /root/gocode/src/github.com/tosone/cplusplus; make depend all"
