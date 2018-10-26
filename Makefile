help:
	@echo "Makefile help"

mahjong-table-go:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: mahjong-table-go
	docker build -t mingz2013/mahjong-table-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/mahjong-table-go


run:
	docker run --net="host" -it mingz2013/mahjong-table-go


.PYONY: help, commit-docker, docker-image, mahjong-table-go, run

