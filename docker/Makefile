app_name                := cplusplus-build
docker_name             := $(app_name)
docker_tag              := dev
docker_container        := $(app_name)

upgrade:
	docker pull alpine:edge

.PHONY: build
build:
	docker build -t $(docker_name):$(docker_tag) .

run:
	if [ ! 'docker inspect -f {{.State.Running}} $(docker_container)' ]; then \
		docker stop $(docker_container); \
	fi
	if [ ! 'docker ps -a | grep $(docker_container))' ]; then \
		docker rm $(docker_container); \
	fi
	docker run -it --hostname $(app_name) --name $(docker_container) --restart always -v $(shell pwd)/..:/root/data -d $(docker_name):$(docker_tag)

exec:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish

exec-build:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish -c "mkdir -p /data/gocode/src; cd /root/data; cp -r vendor/* /data/gocode/src; make depend all"
