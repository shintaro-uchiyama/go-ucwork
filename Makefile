build-local:
	docker build -f ./build/Dockerfile -t go-ucwork .
	docker run -d -v $(PWD):/go/src/app/ --name go-ucwork go-ucwork
run-local:
	docker exec -it go-ucwork ./hello
delete-local:
	docker stop go-ucwork
	docker rm go-ucwork
