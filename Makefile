.PHONY: deps clean build

clean: 
	rm -rf ./example-function/example-function
	
build:
	GOOS=linux GOARCH=amd64 go build -o example-function/example-function ./example-function
	docker build -t npmbin .
	docker run npmbin cat /tmp/npm-layer.zip > npm-layer.zip && unzip npm-layer.zip -d npm-layer

local-run:
	sam local start-api