# gcrypto
[![tests](https://img.shields.io/github/workflow/status/hashamali/gcrypto/tests?label=tests&style=flat-square)](https://github.com/hashamali/gcrypto/actions?query=workflow%3Atests)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gcrypto/security?label=security&style=flat-square)](https://github.com/hashamali/gcrypto/actions?query=workflow%3Asecurity)
[![coverage](https://badgen.net/codecov/c/github/hashamali/gcrypto)](https://codecov.io/gh/hashamali/gcrypto)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gcrypto)](https://goreportcard.com/report/github.com/hashamali/gcrypto)
[![license](https://badgen.net/github/license/hashamali/gcrypto)](https://opensource.org/licenses/MIT)

Thin wrapper around Go's standard crypto library.

#### Methods

* `HMACSHA256`: Generates an HMAC SHA256 of given body with provided key.
* `SHA256`: Generates SHA256 of given body.
* `MD5`: Generates MD5 of given body.

#### Testing

`make test`
