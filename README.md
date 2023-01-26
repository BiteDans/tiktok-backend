# Intro

Tiktok-backend is the backend module for Bytedance Bootcamp Tiktok project. We use Go, Hertz, Gorm and Docker to build this server.

# Running the porject

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
