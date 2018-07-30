[![Build Status](https://travis-ci.org/babbel/codeclimate-go.svg?branch=master)](https://travis-ci.org/babbel/codeclimate-go)

# codeclimate-go

A client written in [go](https://golang.org/) for provisioning [Codeclimate](https://codeclimate.com)

### Usage

```go
go get -u "github.com/babbel/codeclimate-go/codeclimate"
```

```go
import "github.com/babbel/codeclimate-go/codeclimate"
```

### Init new client

```
client, err := codeclimate.NewClient("your_api_key")
```

### Running the tests

```bash
go test -race -v ./...
```

In order to run tests, please, set up environment variable `CODECLIMATE_API_TOKEN` to connect to Codeclimate.

### Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

### Authors

* **Mikhail Chinkov** - *Initial work* - [cazorla19](https://github.com/parabolic)

### License

This project is licensed under the Mozilla License - see the LICENSE.md file for details.
