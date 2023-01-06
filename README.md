# My Base Line 
## _Go Base Line By [AbdullahPrasetio](https://github.com/abdullahPrasetio)_

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/abdullahPrasetio)

Base ini digunakan untuk menjalankan service yang menghasilkan Rest API

- Rest API
- Validation Rest API
- ✨Magic ✨ Code

## Features

- Custom Validation with go playground validator.
- Custom Logger request and response api.

## Tech

My Base Go use tech:

- [GIN](https://github.com/gin-gonic/gin) - Awesome Web Framework!
- [Validator](https://github.com/go-playground/validator) - Package validator implements value validations for structs and individual fields based on tags.
- [Formatter Validator](https://github.com/abdullahPrasetio/validation-formatter) - Awesome custom validator.


## Installation

My Base go requires [Go](https://go.dev/) v1.18+ to run.

Install the dependencies and devDependencies and start the server.

```sh
cd my base
go run main.go
```


## Docker

My Base is very easy to install and deploy in a Docker container.

By default, the Docker will expose port 8025, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
cd app
docker build -t <youruser>/micro-${name}:${version} .
```


## License

MIT

**Free Software, Hell Yeah!**
