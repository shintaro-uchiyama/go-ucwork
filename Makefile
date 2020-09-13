.PHONY: build
build:
	docker build -f ./build/Dockerfile -t go-ucwork .
	docker stop go-ucwork
	docker rm go-ucwork
	docker run -itd -v $(PWD):/go/src/app/ --name go-ucwork go-ucwork
.PHONY: create-user
create-user:
	docker exec -it go-ucwork go run ./cmd/user/createUser.go ${email} ${password}
.PHONY: delete
delete:
	docker stop go-ucwork
	docker rm go-ucwork
