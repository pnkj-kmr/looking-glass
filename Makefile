
release:
	goreleaser release --snapshot --clean 

build:
	goreleaser build --single-target --snapshot --clean

b:
	go build -o lookingglass main.go 

386_b:
	env GOOS=linux GOARCH=386 go build -o lookingglass main.go 

run:
	go run main.go -debug

