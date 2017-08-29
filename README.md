# open311-backend

This is a backend to create reports from the [Smart Cities](https://gitlab.priv.als/witcher/Smart_Cities/commits/open-source) application

## Getting Started

To start using this project first you must install [Golang] (https://golang.org/dl/). 


### Prerequisites

The following dependencies are used in this project:

* [Gin-Gonic](https://github.com/gin-gonic/gin) - HTTP Web Framework 
* [Viper](https://github.com/spf13/viper) - Library used for giving a configuration solution
* [Elasticsearch for go](https://github.com/olivere/elastic) - Elasticsearch functions for Go.
* [Minio-sdk](https://docs.minio.io/docs/golang-client-quickstart-guide) - Minio sdk for Go
* [AWS S3 Go sdk](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html) - AWS S3 sdk for Go

This backend uses a [Elasticsearch](https://www.elastic.co/) for storing reports created in the Smart Cities application, also a [Amazon Web Service S3](https://aws.amazon.com/es/s3/)

for image storage, as a second option for the image storage you could use a [Minio](https://www.minio.io/) these are all required by this backend.

in the [config.toml](config.toml) file you can configure the address of the different services That this backend makes use of.
```toml
[311]
host = "<PUT YOUR AWS S3 ADDRESS HERE>"
[minio]
host = "<PUT YOUR MINIO ADDRESS HERE>"
accessKeyID = "<PUT YOUR ACCESS KEY ID HERE>"
secretAccessKey = "<PUT YOUR SECRET ACCESS KEY HERE>"
[server]
port = 9090
[s3]
region = "<PUT YOUR AWS S3 REGION HERE>"
```


### Installing

Once Golang is installed clone this repository and run the following command at your project root
```
go get
```
This will download all the dependencies required. 


## Deployment

Once you have installed [Elasticsearch](https://www.elastic.co/downloads/elasticsearch) and [Minio](https://www.minio.io/downloads/#minio-server) (or [Amazon Web Service S3](https://aws.amazon.com/es/s3/) configured) run them on your machine.

Go to the project's root and execute the following command in your terminal:

```
go run main.go
```

You could use [Postman](https://www.getpostman.com/postman) for testing all endpoints of this backend.

## Contributing

...

## Versioning

This is the initial version We use [SemVer](http://semver.org/) for versioning.

## Authors

* [Altus Consulting Software Team] (https://github.com/AltusConsulting)

## License

...