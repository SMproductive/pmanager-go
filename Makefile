winFlags = CC=x86_64-w64-mingw32-gcc GOOS=windows CGO_ENABLED=1
linuxFlags = GOOS=linux

default: windows linux

windows:
	${winFlags} go build .
	${winFlags} fyne package -os windows

linux:
	${linuxFlags} go build .
	${linuxFlags} fyne package -os linux

clean:
	go clean
	rm pmanager-go.tar.xz
