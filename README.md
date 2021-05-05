### Monorepo Movie Service

![alt text](https://i.imgur.com/4KIuy9Q.png "System Architecture")
This is how movie service designed. There are 3 gRPC servers which have a specific purpose. Movie service for search and
movie watchlist. User service for User management and auth service for authentication. These 3 services will get bound
with gateway. Gateway will check for access authentication then if it authenticated it will go to designated service.

### Folder Structure

This project will use golang as the main language. Every service will put inside golang/cmd. pkg folder can be accessed
from all services. inside pkg there are many helpers for application functionality as mongo db connection, mongo index.
error handling, grpc helper, http helper, config, crypt, ulid and validator

### Prerequisite

You need to have this following for running this repo:

1. Docker
2. If you're using windows, you need to run wsl2
3. Internet connection
4. Git

#### How to build and run

To run this repo follow this command

- If you don't want to use the zip file clone this repo with

```bash
$ git clone git@github.com:hidayatullahap/go-monorepo-example.git
```

- cd into go-monorepo-example
- run the application with

```bash
$ make start
```

wait for the build, after it finished you can start looking for the available rest API.

Export Monorepo collection into your postaman collection

[Postman download](https://drive.google.com/file/d/1grEUeXn2Bc56MJ3-Ssz8IZaoA0Jy1K7I/view?usp=sharing)
Remember to set postman env

- token: user token
- host: localhost:9094 (default host)

you connect to mongo db in docker with this creds

```
mongodb://hidayatullahagung:12345678@localhost:27020/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false
```

you can clean and stop this repo with

```bash
make clean
```

or to stop with

```bash
make stop
```
