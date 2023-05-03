## Start server
Start kafka, zookeeper and mysql server

``` bash
docker-compose up -d
```

## Create table

``` bash
# go to migration dir
cd migration

# run
go run main.go
```

## Start producer server
``` bash
# go to producer dir
cd producer

# run
go run main.go
```




