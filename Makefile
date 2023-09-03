run-test:
	go test -v ./test
run-app:
	docker-compose up --remove-orphans --build