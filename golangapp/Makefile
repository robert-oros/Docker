up:
	docker-compose up
down:
	docker-compose down
rmi:
	docker rmi $(docker images -a | grep "^<none>")
	docker rmi golangapp_api:latest

clean:
	make down
	make rmi