import os

os.chdir("..")


def build_desktop():
    os.putenv("CGO_ENABLED", "1")
    os.system("go build -o aioDesk.exe desktop/main.go")


def build_server_win():
    os.putenv("CGO_ENABLED", "0")
    os.putenv("GOARCH", "amd64")
    os.putenv("GOOS", "windows")
    os.system("go build -o aio.exe cmd/aio/aio.go")


def build_server_linux():
    os.putenv("CGO_ENABLED", "0")
    os.putenv("GOARCH", "amd64")
    os.putenv("GOOS", "linux")
    os.system("go build -o aio cmd/aio/aio.go")


if __name__ == '__main__':
    build_server_win()
