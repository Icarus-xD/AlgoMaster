start:
	docker-compose build
	docker-compose up

dep:
	go mod download

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

app_image:
	docker build -t algo-master .