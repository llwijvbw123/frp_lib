export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

#build: frpc-lib-dll

# compile assets into binary file
file:
	rm -rf ./assets/frps/static/*
	rm -rf ./assets/frpc/static/*
	cp -rf ./web/frps/dist/* ./assets/frps/static
	cp -rf ./web/frpc/dist/* ./assets/frpc/static

fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .

vet:
	go vet ./...

frps:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/frps ./cmd/frps

frpc:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/frpc ./cmd/frpc

frpc-lib-static:
	env CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS) -extldflags=-static" -buildmode=c-archive -o bin/ ./cmd/libfrpc

frpc-lib-unix:
	env CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS) -extldflags=-static" -buildmode=c-shared -o bin/ ./cmd/libfrpc

frpc-lib-windows64:
	env CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS) -extldflags=-static" -buildmode=c-shared -o bin/x86_64/frpc.dll ./cmd/libfrpc
	cmd /C dll2lib.bat 64 bin/x86_64/frpc.dll
	mv frpc.lib bin/x86_64/frpc.lib

frpc-lib-windows32:
	export GOARCH=386
	env GOGCCFLAGS=-m32 CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS) -extldflags=-static" -buildmode=c-shared -o bin/x86/frpc.dll ./cmd/libfrpc
	cmd /C dll2lib.bat 32 bin/x86/frpc.dll
	mv frpc.lib bin/x86/frpc.lib

frpc-lib-android:
	go get golang.org/x/mobile/cmd/gobind
	go get golang.org/x/mobile/cmd/gomobile
	gomobile init
	gomobile bind -v -o bin/android/libfrp.aar -target=android ./cmd/libfrpc_mobile

frpc-lib-ios:
	go get golang.org/x/mobile/cmd/gobind
	go get golang.org/x/mobile/cmd/gomobile
	gomobile init
	gomobile bind -v -o bin/ios/ -target=ios ./cmd/libfrpc_mobile

test: gotest

gotest:
	go test -v --cover ./assets/...
	go test -v --cover ./cmd/...
	go test -v --cover ./client/...
	go test -v --cover ./server/...
	go test -v --cover ./pkg/...

e2e:
	./hack/run-e2e.sh

e2e-trace:
	DEBUG=true LOG_LEVEL=trace ./hack/run-e2e.sh

alltest: vet gotest e2e
	
clean:
	rm -f ./bin/frpc
	rm -f ./bin/frps
