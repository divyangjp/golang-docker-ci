### Pre-Requisits
- go
- make
- docker
- docker-compose

### Build and Run locally

Build executable and docker image (otc-app)
```sh
make build
```

Run local binary
```sh
SECRET=<SECRETKEY> make run
```
### Build/Run/Stop using docker-compose locally
```sh
make build-docker-compose
make run-docker-compose
make stop-docker-compose
```
Default value of `SECRET` is set to `DefaultSecret` inside docker-compose.  
To set `SECRET` value, append make command with `SECRETKEY`
```sh
make run-docker-compose SECRETKEY=TOP5ecREt
```

App is exposed onto port **8080**
Local APIs
```sh
http://127.0.0.1:8080/metrics
http://127.0.0.1:8080/health
http://127.0.0.1:8080/token
```

### Docker Build ###
Multistage docker build to keep the docker image size minimum.  
Google Distroless image used to install app. Keeps the security attack surface minimum.

### Build CI
Github Actions are used for CI (https://github.com/divyangjp/golang-docker-ci/actions)   
**master-build-push**:   
_Condition_ - On push to `master` branch AND when app files modified (_go.*_ or _*.go_)   
_Outcome_ - Docker container image built and tagged with Git SHA. Pushed to hub.docker.com at https://hub.docker.com/r/divyangjp/otc-app/tags   

DockerHub login credentials are stored into Github repository secrets   
- DOCKER_HUB_USERNAME
- DOCKER_HUB_PAT

**Pull Request CI**:   
_Condition_ - Pull Request is raised on branches `master` or `releases/**`   
_Outcome_ - Runs `make clean test build` to test for any issues with PR. Additional goodies like Jira number check, linting, code scanning etc can be added to this workflow   

**CODEOWNERS**: To automatically notify code owners when PR is raised to modify certain files/directories   

### Service interface

- POST /token

  Return a token based on a shared secret. The shared secret is passed in the
  environment as the variable `SECRET`.

- GET /health

  Used to check the service is 'up'. It should return an HTTP code >= 200

- GET /metrics

  Return some basic metrics about the running service.

### Secrets
- Only read/load secrets into container at runtime
- Use services like AWS Secrets Manager, GCP Secret Manager or Hashicorp Vault to safely handle secrets
- Use IAM, Service Accounts, WorkloadIdentity etc to access secrets from secret stores
