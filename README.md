# shipsport

Author: Piotr Jasiak

The service streams data from the json file and saves the entities in the repository

# Usage

To start the service:
1. Create docker image:
```sh
docker build . -t shipport:0.1
```
2. Run image:
```sh
docker run -v "${PWD}/ports.json:/ports.json" shipport:0.1 ../ports.json
```

# Tests

To executes tests run in root directory:
```sh
go test ./...
```
