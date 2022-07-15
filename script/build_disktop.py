import os

os.putenv("CGO_ENABLED", "1")

os.system("go build -o aioDesk.exe desktop/main.go")
