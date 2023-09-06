BINARY_NAME=sprapp

build:
	git pull
	go mod vendor
	CGO_ENABLED=1 go build -ldflags "-w -s" app.go

run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

git:
	git add .
	git commit -m "$m"
	git push -u origin main
