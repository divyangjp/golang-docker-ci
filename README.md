Testing PR CI
Try again
## Building the binary

You will need a Go development environment installed to build the binary.
To build the binary execute the following:

```shell

go build

```

## Service interface

- POST /token

  Return a token based on a shared secret. The shared secret is passed in the
  environment as the variable `SECRET`.

- GET /health

  Used to check the service is 'up'. It should return an HTTP code >= 200

- GET /metrics

  Return some basic metrics about the running service.
