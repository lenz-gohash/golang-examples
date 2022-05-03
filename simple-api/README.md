# simple-api

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A simple restful api example with a username and password authentication 

## Table of Contents

- [Usage](#usage)
- [API](#api)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)


## Usage

Run the server 
```
go run main.go
```

Access the unprotected endpoint 
``` 
curl http://localhost:8080
```

Access the protected endpoint 
```
curl -u lenz:gohash  http://localhost:8080/secret
```


## API

Protected endpoint: http://localhost:8080/

Unprotected endpoint: http://localhost:8080/secret


## Maintainers

[@lenz-gohash](https://github.com/lenz-gohash)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2022 Lenz Paul
