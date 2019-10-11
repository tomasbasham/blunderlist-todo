# blunderlist-todo [![Build Status](https://travis-ci.com/tomasbasham/blunderlist-todo.svg?branch=master)](https://travis-ci.com/tomasbasham/blunderlist-todo) [![Maintainability](https://api.codeclimate.com/v1/badges/428e6cae5d8321a778ed/maintainability)](https://codeclimate.com/github/tomasbasham/blunderlist-todo/maintainability)

A fictitious todo application through which to teach how to implement a
microservice architecture. For the full list of services required to run this
application visit
[Blunderlist](https://github.com/tomasbasham?utf8=✓&tab=repositories&q=blunderlist)
on GitHub.

This repository implements an API that manages tasks and their creation,
providing an isolated data abstraction from other components of the overall
system. Data from this service may be further decorated with data obtained from
other services in some downstream closer to the client.

The intent of this repository is to provide the most optimal API surface for
the domain with which it is concerned, issuing unique resource handles that may
be referred to in other services.

The code here attempts to follow the principles of [Domain Driven
Design](https://www.google.com/search?q=domain-driven+design); where my take on
the subject can be seen on [my
blog](https://tomasbasham.dev/development/2019/10/26/domain-driven-design-in-practice).

## Prerequisites

You will need the following things properly installed on your computer.

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)

## Installation

* `git clone <repository-url>` this repository
* `cd blunderlist-todo`
* `docker build -t todo .`

## Running / Development

* `docker run --rm -it -p 8080:8080 --env-file .env todo`
* Visit your app at [http://localhost:8080](http://localhost:8080).

### Rebuilding Protobuf Definitions

When changes to the API surface are made it is necessary to update the
generated protobuf definitions and any code that uses them. A docker image with
the `protoc` compiler has been made to make this far simpler.

* `docker run --rm -v $(pwd):/usr/src/app tomasbasham/protoc:3.6.1-go -I ./proto --go_out=plugins=grpc:. todo.proto`

## Further Reading / Useful Links

* [Go](https://golang.org/)
* [gRPC](https://grpc.io/)
* [protobuf](https://developers.google.com/protocol-buffers/)
