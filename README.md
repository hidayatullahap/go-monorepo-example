### Monorepo Movie Service

![alt text](https://i.imgur.com/4KIuy9Q.png "System Architecture")
This is how movie service designed. There are 3 gRPC servers which have a specific purpose. Movie service for search and
movie watchlist. User service for User management and auth service for authentication. These 3 services will get bound
with gateway. Gateway will check for access authentication then if it authenticated it will go to designated service.

### Prerequisite

You need to have this following for running this repo:

1. Docker
2. If you're using windows, you need to run wsl2
3. Internet connection

#### How to build and run

To run this repo type this command

```bash
make start
```

wait for the build, after it finished you can start looking for the available rest API.

Export Monorepo collection into your postaman collection

[Postman download](https://drive.google.com/file/d/1zsZqPK7-jGQfRyQIQqn38QH2XeUlHO8Y/view?usp=sharing)
Remember to set postman env

- token: user token
- host: localhost:9094 (default host)

you can clean and stop this repo with

```bash
make clean
```

or to stop with

```bash
make stop
```
