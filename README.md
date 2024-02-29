# Go Example Proto

This is an example project mirroring my [example project for prost in rust](https://github.com/krlohnes/prost_example)

## Build requirements
[protoc installation instructions](https://grpc.io/docs/protoc-installation/)

Just in case the link goes dead, here's the instructions

Linux, using apt or apt-get, for example:
```sh
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

MacOS, using Homebrew:
```sh
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

## Building

This build can likely be automated by a Makefile, but I want to show what's involved since go's
tooling can't handle this on its own like rust's can.

### Install protoc-gen-go
```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Make sure the go option is specified in the protobuf files
This can also be handles by a build tool like bazel, but it seems simple enough to include a single
option in your protofile if you're not working in a largely polyglot environment.

See [items.proto](proto/items.proto) for an example

### Compile the protobufs
The following script uses `find` to automate adding the protobuf files to the command.

```sh
$ protoc -I=protos/ \
         --go_out=src/protos \
         $(find protos -iname "*.proto")
```

### Make sure your protobufs aren't checked in

Or check them in. If you do check them in, however, it would begood to have a CI process that can
recompile them and ensure that the generated protos and the existing code for the branch are in sync
or error if not.

### Run the go stuff
```sh
$ go run ./src/main.go
```
