linux64: export GOOS = linux
linux64: export GOARCH = amd64
linux64: export CGO_ENABLED=1

linux32: export GOOS = linux
linux32: export GOARCH = 386
linux32: export CGO_ENABLED=1

windows64: export GOOS = windows
windows64: export GOARCH = amd64
windows64: export CGO_ENABLED=1
windows64: export CXX = x86_64-w64-mingw32-g++
windows64: export CC = x86_64-w64-mingw32-gcc

windows32: export GOOS = windows
windows32: export GOARCH = 386
windows32: export CGO_ENABLED=1
windows32: export CXX = i686-w64-mingw32-g++
windows32: export CC = i686-w64-mingw32-gcc

all: linux32 linux64 windows32 windows64


linux64:
	go build -o mouseMoverLinux64

linux32:
	go build -o mouseMoverLinux32

windows32:
	go build -ldflags -H=windowsgui -o mouseMoverWindows32.exe

windows64:
	go build -ldflags -H=windowsgui -o mouseMoverWindows64.exe