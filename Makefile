run:
	go run main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o marine-traffic-linux main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o marine-traffic-mac main.go

docker-build:
	docker build -t marine-traffic .

docker-run:
	docker run -p 80:80 marine-traffic