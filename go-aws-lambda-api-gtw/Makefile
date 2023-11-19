lambda: build zip

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main 

zip:
	zip main main