# cyberuranus-interview

> Only for Cyberuranus Interview

## Build

```bash
go build .
```

## Run

```bash
go run main.go
```

## Run with docker

```bash
docker build -t cu-server .
docker pull mysql
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
docker run --link mysql -p 8080:8080 cu-server
```