# Intro

Tiktok-backend is the backend module for Bytedance Bootcamp Tiktok project. We use Go, Hertz, Gorm and Docker to build this server.

# Running the project

To run the project, we first run the following command to fetch the dependencies:

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
To add environment variables, create `/pkg/configs/.env` and add variables into the file. Note that the file is ignored by git.

For example
```
SECRET_1=SECRET1
SECRET_2=SECRET2
JWT_SECRET=THISISASECRET
```

To get env variables, use `os.Getenv("SECRET_NAME")`, for example
```go
var key = os.Getenv("JWT_SECRET") // returns "THISISASECRET"
```