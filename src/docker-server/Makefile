export tag=v1.0
build:
	echo "building httpserver binary"
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/httpserver .

image: build
	echo "building httpserver image"
	docker build -t httpserver:${tag} .

push: image
	echo "pushing httpserver:${tag} to aliyun repository"
	docker tag httpserver:${tag} registry.cn-chengdu.aliyuncs.com/golang-project1/httpserver:${tag}
	docker push registry.cn-chengdu.aliyuncs.com/golang-project1/httpserver:${tag}
rundocker: push
	docker run --rm --name=http-server -d -p 8080:8080 registry.cn-chengdu.aliyuncs.com/golang-project1/httpserver:${tag}
	
