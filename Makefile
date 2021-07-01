.PHONY: run clean build

run:
	docker-compose up -d bookshelf

build:
	docker-compose build --no-cache

clean:
	docker-compose down --rmi all
	docker image prune -f