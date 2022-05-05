# proto-go-course

Protobuf Go 학습 archiving 목적

[![build main branch](https://github.com/gain620/proto-go-course/actions/workflows/build.yml/badge.svg)](https://github.com/gain620/proto-go-course/actions/workflows/build.yml) [![Lint protobuf](https://github.com/gain620/proto-go-course/actions/workflows/lint.yml/badge.svg)](https://github.com/gain620/proto-go-course/actions/workflows/lint.yml)

## Notes

### `Windows`

- I recommend you use powershell (try to update: [see](https://github.com/PowerShell/PowerShell/releases)) for following this course.
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install))


### Build

#### `Linux/MacOS`

```shell
go mod tidy
make
```

#### `Windows - Chocolatey`
```shell
choco install make
go mod tidy
make
```

#### `Windows - Without Chocolatey`

```shell
protoc -Iproto --go_opt=module=github.com/Clement-Jean/proto-go-course --go_out=. proto/*.proto

go mod tidy
go build -o proto-go-course.exe .
```

## Run

```
./proto-go-course
```

or

```
./proto-go-course.exe
```