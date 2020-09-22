.PHONY: build
build:
	docker build -f ./build/Dockerfile -t go-ucwork .
	docker stop go-ucwork
	docker rm go-ucwork
	docker run -itd -v $(PWD):/go/src/app/ -p 8080:8080 --name go-ucwork go-ucwork
.PHONY: run-server
run-server:
	docker exec -it go-ucwork reflex -r "(.go|go.mod)" -s go run pkg/server.go
.PHONY: create-user
create-user:
	docker exec -it go-ucwork go run ./cmd/user/create_user.go ${email} ${password}
.PHONY: get
get:
	docker exec -it go-ucwork go get ${dep}
.PHONY: delete
delete:
	docker stop go-ucwork
	docker rm go-ucwork
