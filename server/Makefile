build_linux:
	env GOOS=linux GOARCH=386 go build lg.go

build:
	go build -o lookingglass main.go && mv lookingglass ../release/ && cp -rf conf ../release/

386_build:
	env GOOS=linux GOARCH=386 go build -o lookingglass main.go && mv lookingglass ../release/ && cp -rf conf ../release/

run:
	go run main.go 