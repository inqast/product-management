build-all:
	cd checkout && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build
	cd notifications && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	sudo docker compose up --force-recreate --build
	cd loms && make migrate-up
	cd checkout && make migrate-up
	cd notifications && make migrate-up

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

generate-all:
	cd checkout && make generate
	cd loms && make generate