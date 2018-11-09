help:
	@echo "Makefile help"

clean:
	rm mahjong-table-go


mahjong-table-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: mahjong-table-go
	docker build -t mingz2013/mahjong-table-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/mahjong-table-go

run:
	docker run -d --link redis-mq:redis-mq --name mahjong-table-go -it mingz2013/mahjong-table-go:latest



.PYONY: help, commit-docker, docker-image, mahjong-table-go, run

