Sqooba's common [Go](https://golang.org) packages, a.k.a github.com/sqooba/go-common
====

| Package | Main usage |
| ------- | ---------- |
| [data-structure](./data-structure) | Helper function around maps |
| [env-parsing](./env-parsing) | [flag](https://golang.org/pkg/flag/) and [envconfig](https://github.com/kelseyhightower/envconfig) extensions |
| [healthchecks](./healthchecks) | Tooling around [docker healthcheck](https://github.com/docker/distribution/health) |
| [logging](./logging) | Tooling around [logrus](https://github.com/sirupsen/logrus) |
| [random](./random) | Some random data and password generators |
| [time](./time) | [time](https://golang.org/pkg/time/) and [Duration](https://golang.org/pkg/time/#Duration) parser and formatter |
| [version](./version) | Version inclusion at build time, for instance `-ldflags "-X version.GitCommit=${GIT_COMMIT}"` |
