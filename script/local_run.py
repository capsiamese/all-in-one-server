import os
import subprocess
import webbrowser

os.chdir("..")
# os.putenv("GOARCH", "amd64")
# os.putenv("GOOS", "linux")
# os.putenv("CGO_ENABLED", "0")

os.system("ent generate ./ent/schema")
os.system("swag fmt")
os.system("swag init -g internal/controller/http/v1/route.go -parseInternal")

dsn = os.getenv("ENT_MIGRATE")
os.system(f"go run cmd/aio/aio.go mig -e -d {dsn}")

proc = subprocess.Popen(["go", "run", "cmd/aio/aio.go"],
                        stderr=subprocess.PIPE,
                        stdout=subprocess.PIPE,
                        encoding='utf-8')

webbrowser.open('http://localhost:8080/swagger/index.html')

while True:
    line = proc.stdout.readline()
    if line:
        print(line, end='')
