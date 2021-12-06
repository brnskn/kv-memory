# Easy to Use In-Memory Key-Value Store

A project that provides an in-memory key-value store as a REST API. Also, it's containerized and can be used as a microservice.

## Installation

Run these commands on this project's root:

```bash
docker build -t "memory" .
docker run -p 127.0.0.1:8080:8080 -d memory
```

There are some environment variables that this project uses, and they can be set with docker run

```bash
docker run -p 127.0.0.1:8081:8081 -d \
-e AUTO_SAVE_INTERVAL=10 \
-e PORT=8081 \
memory
```

## Demonstration

This project is deployed to Heroku: <https://thawing-fjord-72264.herokuapp.com/>

## API Documentation

The API documentation: <https://documenter.getpostman.com/view/551409/UVJhEFaU>

## GODOC Documentation

To see the godoc documentation run these commands on the project's root

```bash
godoc -http=:6060 -play
```

And then visit this link: <http://127.0.0.1:6060/pkg/github.com/brnskn/kv-memory/>

## Roadmap

- [x] Use repository and singleton design patterns.
- [x] Write a readme.
- [x] Add comment lines for go doc.
- [x] Add postman API doc link.
- [x] Write tests.
- [x] Add a Dockerfile.
- [x] Deploy to Heroku.
- [x] Add a LICENSE.
- [x] Add GitHub actions(test, lint, build).
- [ ] Add HTTP request logs
