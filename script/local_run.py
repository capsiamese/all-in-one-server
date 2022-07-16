import os
import subprocess
import webbrowser

os.chdir("..")

os.putenv("CGO_ENABLED", "0")


def code_gen():
    os.system("ent generate ./ent/schema")
    os.system("swag fmt")
    os.system("swag init -g internal/controller/http/v1/route.go -parseInternal")


def migrate():
    dsn = os.getenv("ENT_MIGRATE")
    os.system(f"go run cmd/aio/aio.go mig -e -d {dsn}")


def proto_gen():
    for entry in os.scandir("proto"):
        if entry.is_file() and entry.name.endswith(".proto"):
            os.system(f"protoc --go_out=. --proto_path=proto proto/{entry.name}")


if __name__ == '__main__':
    proto_gen()
    exit(0)
    code_gen()
    migrate()
    proc = subprocess.Popen(["go", "run", "cmd/aio/aio.go"],
                            stderr=subprocess.PIPE,
                            stdout=subprocess.PIPE,
                            encoding='utf-8')

    webbrowser.open('http://localhost:8080/swagger/index.html')

    while True:
        line = proc.stdout.readline()
        if line:
            print(line, end='')
