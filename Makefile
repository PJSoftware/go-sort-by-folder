BINARY=bin/SortByFolder.exe

.PHONY: build run install clean

build:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY} main.go

run: ${BINARY}
	${BINARY}

install: ${BINARY}
	cp ${BINARY} /h/_rip/_series

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

${BINARY}: build
