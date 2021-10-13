.PHONY: up down test

up:
	cd test && docker-compose up -d --build

down:
	cd test && docker-compose down

test:
	go test ./... -v

test_local:
	BITBUCKET_BASE_URL='http://localhost:5000' \
	BITBUCKET_USERNAME=test \
	BITBUCKET_PASSWORD=test \
	make up test down

test_integration:
	go test ./... -v -tags integration