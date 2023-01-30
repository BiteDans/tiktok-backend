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

# User Authentication

We have created a middleware `AuthMiddleware` for simple token authentication.

Use `GET /douyin/user` as an example.

First, find the corresponding function that returns the middelwares for the route
```go
// /biz/router/douyin/core/user.go
_douyin.GET("/user", append(_userMw(), core.UserInfo)...)
```
Then put the middleware in the returning slice
```go
// /biz/router/douyin/core/middleware.go
func _userinfoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.AuthMiddleware(),
	}
}
```
