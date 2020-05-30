.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculator-grpc calculator/calculator-grpc/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculator-client calculator/calculator-client/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
