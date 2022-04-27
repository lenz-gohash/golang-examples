# simple-tcp

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

An example of a simple TCP server and client.

The client sends a message and the server prints it to the console.

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

In one terminal, run the server 
```
./simple-tcp server
```

In another terminal, run the client
```
./simple-tcp client
```

Output:

```
2022/04/27 10:57:38 Received connection from 127.0.0.1:49192
2022/04/27 10:57:38 Received 13 bytes: Hello! World!
```


## Maintainers

[@lenz-gohash](https://github.com/lenz-gohash)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2022 Lenz Paul
