
REVISION := `git rev-list -1 --abbrev-commit HEAD`
BUILD_DATE := `date`
BUILD_CONTEXT := `whoami`"@"`hostname`
VERSION := "master"

clean:
	rm -rvf build/ || true
	rm -rf app-ui || true

mkdirs:
	mkdir build || true
	mkdir app-ui || true

install:
	go get -v github.com/GeertJohan/go.rice/rice
	cd ui && npm install
	go get -v

build_frontend: mkdirs
	cd ui && npx ng build --prod --output-path ../app-ui/

embed: build_frontend
	rm -v rice-box.go || true
	rice embed-go

build: embed
	cd cmd && go build -v -ldflags "-X main.Revision=${REVISION} -X \"main.Version=${VERSION}\" -X \"main.BuildTimestamp=${BUILD_DATE}\" -X main.BuildContext=${BUILD_CONTEXT}" -o ../build/ogwc


