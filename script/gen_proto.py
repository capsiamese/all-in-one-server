import os

for entry in os.scandir("../proto"):
    if entry.is_file() and entry.name.endswith(".proto"):
        os.system(f"protoc --go_out=.. --proto_path=../proto ../proto/{entry.name}")
