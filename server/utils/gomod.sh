.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/dockBattery dockBattery/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/dockStatus dockStatus/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/stationStatus stationStatus/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: build
	sls deploy --verbose --stage $(STAGE)

gomodgen:
	#sudo chmod u+x gomod.sh
	./gomod.sh