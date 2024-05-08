up: ### Run docker
	docker run

down: ### Down docker
	docker down

linter-golangci: ### check by golangci linter
	golangci-lint run

test: ### run test
	go test -v ./...

coverage-html: ### run test with coverage and open html report
	go test -coverprofile=cvr.out ./...
	go tool cover -html=cvr.out
	rm cvr.out

coverage: ### run test with coverage
	go test -coverprofile=cvr.out ./...
	go tool cover -func=cvr.out
	rm cvr.out
