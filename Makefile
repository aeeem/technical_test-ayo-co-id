all: build


dockerbuild:
	swag init
	go build 
	docker build --platform linux/amd64 -f DOCKERFILE . -t tech-test
	docker compose -f docker-compose.yaml up --build -d
	docker logs -f tech-test