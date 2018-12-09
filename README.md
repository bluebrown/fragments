# Cloud-Native GraphQL API

This is my humble try to implement a GraphQL backend. Here I will document what I find out and how I set things up.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

This guide assumes you have a working go installation and docker/docker-compose setup. 

Install godep:
```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```
navigate into gopath/src/github.com/USERNAME/, clone this repository and enter the project:
```
$ cd ~/go/src/github.com/bluebrown
$ git clone https://github.com/bluebrown/fragments.git &&cd fragments
```

### Installing

A step by step series of examples that tell you how to get a development env running.

First make sure the depenedencies are installed using godep:
```
$ dep ensure
```

Build the Docker image:
```
$ docker-compose build
```

Spin up the container and run the go-code:
```
$ docker-compose run --service-ports api /bin/bash
$ go run .
```

Request all message-texts from GraphQL-API:
```
$ curl -XPOST -d '{"query": "{ messages{ text } }"  }' localhost:8080/api -w "\n"
```

Do more complex query to get a users name and mail by ID:
```
$ curl -XPOST -d '{"query": "{ user( ID: \"1\" ){ name mail } }"  }' localhost:8080/api -w "\n"
```

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Development

How to use for development

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Go](http://www.dropwizard.io/1.0.2/docs/) - Language
* [graph-gophers](https://github.com/graph-gophers/graphql-go) - Framework
* [GoDep](https://rometools.github.io/rome/) - Dependencies
* [Docker](https://rometools.github.io/rome/) - DevOps

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
