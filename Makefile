BINARY=SortByFolder.exe

build:
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY} main.go

run: bin/${BINARY}
	bin/${BINARY}

install: bin/${BINARY}
	cp bin/${BINARY} /h/_rip/_series

clean:
	if [ -f bin/${BINARY} ]; then rm bin/${BINARY}; fi

bin/${BINARY}: build
