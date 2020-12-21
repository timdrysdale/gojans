![alt text][logo]

A golang client adapter for [Janssen Project](https://github.com/JanssenProject/home) services, starting with the [jans-auth-server](https://github.com/JanssenProject/jans-auth-server)

[Go, golang, Janssen, Gluu, oxAuth, jans-auth-server]

[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#experimental) [![Go Reference](https://pkg.go.dev/badge/github.com/timdrysdale/gojans.svg)](https://pkg.go.dev/github.com/timdrysdale/gojans) ![alt text][report] [![GitHub license](https://img.shields.io/github/license/timdrysdale/gojans)](https://github.com/timdrysdale/gojans/blob/main/LICENSE) 

For questions either raise an issue or visit the [gopher-slack](https://invite.slack.golangbridge.org/) channel [#gojans](https://gophers.slack.com/app_redirect?channel=gojans)

## Background

[Janssen](https://github.com/JanssenProject/home) (`jans` for short) is a [Linux Foundation](https://www.linuxfoundation.org/) project providing a cloud-native identity and access management (IAM) platform, building on code developed and donated by [Gluu](https://www.gluu.org/). If you are still considering possible alternatives, then you should read [this opinion](https://www.gluu.org/blog/keycloak-is-the-next-centos/) about what Redhat dropping CENTOS means for its sibling [keycloak](https://www.keycloak.org/). 

## About

This adapter aims to ease the task of building golang clients which interact with `jans`, in much the same way that [gocloak](https:github.com/Nerzal/gocloak) does for keycloak. A key difference is that `gojans` is auto-generated, whereas `gocloak` is crafted by hand. The steps to create `gojans` require some human intervention, so this repo exists to save you the hassle. The procedure is

 - start with `openapi3.0` spec for Gluu oxAuth server 4.2, 
 - convert to `openapi2.0/swagger2.0` using [apimatic.io](https://apimatic.io)
 - modify by hand to clean up the naming and some type defintions
 - generate client code with [go-swagger](https://github.com/go-swagger/go-swagger) 


## Usage

To use the client code in your project, import

```
import (
	jac "github.com/timdrysdale/gojans/pkg/jansAuth/client"
	jam "github.com/timdrysdale/gojans/pkg/jansAuth/models"
)
```

## Project Organisation

The `swagger` files are in `./api/openapi-spec`. The client adapter code is in `.pkg/jansAuth`. The project directory tree is:

```
.
├── api
│   └── openapi-spec
├── img
├── internal
│   └── curl
└── pkg
    └── jansAuth
        ├── client
        └── models
```

## CI

You can pre-run the linting with 

```
golangci-lint run ./...  
```

Using `act` may fail on your local machine due to cgo dependencies.


## Mocking / Testing

Test with your own instantiation of the auth-server, or consider mocking (perhaps automatically with [apisprout](https://github.com/danielgtaylor/apisprout) although it may require the examples to be updated for more complete testing).

[logo]: ./img/logo.png "gopher grabs a ride with the janssen racing pigeon"
[report]: ./img/report.svg "badge from https://goreportcard.com/"


