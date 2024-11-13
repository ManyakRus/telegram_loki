SERVICENAME=telegram_loki
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./cmd/$(SERVICENAME)/main.go
FILEAPP=./bin/$(SERVICENAME)

NEW_REPO=github.com/ManyakRus/telegram_loki


run:
	clear
	./make_version.sh
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	./bin/$(SERVICENAME)
mod:
	clear
	go get -u ./...
	go mod tidy -compat=1.22
	go mod vendor
	./make_version.sh
	go fmt ./...
build:
	clear
	./make_version.sh
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cp $(FILEAPP) "./bin_prod/$(SERVICENAME)"
	cd ./cmd && \
	./VersionToFile.py

lint:
	clear
	./make_version.sh
	go fmt ./...
	golangci-lint run ./internal/...
	golangci-lint run ./pkg/...
	gocyclo -over 10 ./internal
	gocyclo -over 10 ./pkg
	gocritic check ./internal/...
	gocritic check ./pkg/...
	staticcheck ./internal/...
	staticcheck ./pkg/...
run.test:
	clear
	./make_version.sh
	go fmt ./...
	go test -coverprofile cover.out ./internal/v0/app/...
	go tool cover -func=cover.out
newrepo:
	clear
	./make_version.sh
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml
conn:
	clear
	image_connections ./internal docs/connections.graphml $(SERVICENAME)
lines:
	clear
	./make_version.sh
	go_lines_count ./ ./docs/lines_count.txt 10
licenses:
	golicense -out-xlsx=./docs/licenses.xlsx $(FILEAPP)
gocyclo:
	golangci-lint run ./... --disable-all -E gocyclo -v
