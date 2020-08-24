build-image:
	docker build -t migueloli/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

lint:
	$$GOBIN/golint ./...
	go fmt -n ./...

unit-test:
	go test ./...