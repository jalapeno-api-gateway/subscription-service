# push-service
The push-service is part of the Jalapeño API Gateway. It allows SR-Apps to subscribe to collections and counters.

## gRPC
- When the file `proto/push-service/push-service.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/push-service --go_out=./proto/push-service --go_opt=paths=source_relative --go-grpc_out=./proto/push-service --go-grpc_opt=paths=source_relative ./proto/push-service/push-service.proto
```
- When the file `proto/tsdb-feeder/tsdb-feeder.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/tsdb-feeder --go_out=./proto/tsdb-feeder --go_opt=paths=source_relative --go-grpc_out=./proto/tsdb-feeder --go-grpc_opt=paths=source_relative ./proto/tsdb-feeder/tsdb-feeder.proto
```
- Replace content of okteto.yml with the following:
```yml
name: push-service
autocreate: true
image: okteto/golang:1
command: bash
namespace: jagw-dev-<namespace-name>
securityContext:
  capabilities:
    add:
      - SYS_PTRACE
volumes:
  - /go/pkg/
  - /root/.cache/go-build/
  - /root/.vscode-server
  - /go/bin/
  - /bin/protoc/
sync:
  - .:/usr/src/app
forward:
  - 2350:2345
  - 8085:8080
environment:
  - GRAPH_DB_FEEDER_ADDRESS=gdbf-svc:9000
  - TSDB_FEEDER_ADDRESS=tsdb-service:9000
```
