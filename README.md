# Intro

Tiktok-backend is the backend module for Bytedance Bootcamp Tiktok project. We use Go, Hertz, Gorm and Docker to build this server, GCP to host the project and AWS S3 to store the files.

The backend (production) is hosted at http://35.206.100.23:8888 on Google Cloud Platform. Feel free to test it out.

# Running the project

Before running the project, make sure you have done the following steps:

1. Download ffmpeg at https://ffbinaries.com/downloads.
2. At ~/.aws, create a file named credentials, and type in

<code>
aws_access_key_id = ***
aws_secret_access_key = ***
</code>

Then, to run the project, we first run the following command to fetch the dependencies:

```console
go mod tidy
```

Then, we run

```console
docker-compose up
```

to start the MySQL service.

Once done, run

```console
go build
./tiktok-backend.exe
```

to start the server.

To connect the database from your local machine, run

```console
mysql -h 127.0.0.1 -P 3307 -u gorm -p
// password=gorm
```

# IDL(Interface Description Language)

We use the command line tool `hz` provided by `hertz` to generate the basic code.

`hz` looks at your interface definitaions in `.thrift` files (e.g. /idl/hello.thrift) and automatically generates a skaffold.

## Start a project

After creating a `.thrift` file

```console
// under GOPATH
hz new -idl idl/hello.thrift

// outside of GOPATH
hz new -idl idl/hello.thrift -mod <module name>
```

## Update a project

If we wish to continue using `hz` to generate code in the future, we need to update our .thrift files to include new interface definitions and run

```console
hz update -idl idl/hello.thrift
```

For more details, refer to [hz's official documents](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/).

# Environment Variables

Environment variables and secrets are stored in `.env` and `.secrets`.

To add variables:

```
MY_ENV_VAR=THIS_IS_AN_ENV_VAR
JWT_SECRET=THIS_IS_A_SECRET
```

To get env variables, use `os.Getenv("SECRET_NAME")`, for example

```go
var key = os.Getenv("JWT_SECRET") // returns "THIS_IS_A_SECRET"
```
