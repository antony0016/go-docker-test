## GO-TEST

#### Docker image build:

```shell
docker build -t go-test .
```

#### Create a docker container :

##### Docker run:

run

```shell
docker run -d --name go-test -p 3000:3000 go-test
```

stop and remove container

```shell
docker stop go-test
docker rm go-test
```

[comment]: <> (> if you want to connect to local postgres server change environment variable **DB_HOST** to **host.docker.internal** in **.env** file)

##### Docker compose up:

run

```shell
docker-compose up -d
```

stop and remove container

```shell
docker-compose down -d
```
