# curl-docker

Small program to easily curl from Docker (in particular on Windows via an npipe).

## Build

```
go build
```

## Usage

```
.\curl-docker.exe -pipe \\.\pipe\dockerBackendApiServer -url http://localhost/versions
```
