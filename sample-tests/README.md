# sample-tests

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

Examples of Go unit tests

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Install

```
go build
```

## Usage

Run main function: 

```
./sample-tests
```


To run all tests: 
```
go test ./...
``` 

**note**: Some tests are purposely failing. This is to show the output of the
failing tests.

<br>

Generating coverage report: 

```
go test ./... -coverprofile=c.out
```

Visualizing coverage report by function: 

```
go tool cover -func=c.out
```

View coverage report in browser: 

```
go tool cover -html=c.out
```


## Maintainers

[@lenz-gohash](https://github.com/lenz-gohash)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2022 Lenz Paul
