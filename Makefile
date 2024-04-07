.PHONY: build clean zip deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	echo "Building lambada binaries"
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o build/saveDevice/bootstrap saveDevice/main.go
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o build/sendSquawk/bootstrap sendSquawk/main.go

zip:
	zip -j build/saveDevice.zip build/saveDevice/bootstrap
	zip -j build/sendSquawk.zip build/sendSquawk/bootstrap squawker-5b498-firebase-adminsdk-wdxa6-9b77127a4c.json

clean:
	rm -rf ./build ./vendor

deploy: clean build zip
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh