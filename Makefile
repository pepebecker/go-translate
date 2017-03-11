build:
	go build

osx:
	GOOS=darwin go build

linux:
	GOOS=linux go build

windows:
	GOOS=windows go build

run:
	./project1