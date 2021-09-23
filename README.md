# push-service
The push-service is part of the Jalapeño API Gateway. It allows SR-Apps to subscribe to collections and counters.

## gRPC
- When the file `proto/subscriptionservice/subscriptionservice.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/subscriptionservice --go_out=./proto/subscriptionservice --go_opt=paths=source_relative --go-grpc_out=./proto/subscriptionservice --go-grpc_opt=paths=source_relative ./proto/subscriptionservice/subscriptionservice.proto
```
## Setting Up Development Environment
Make sure you have setup the [global development environment](https://gitlab.ost.ch/ins/jalapeno-api/request-service/-/wikis/Development-Environment) first.

## Initialize Okteto
```bash
$ git clone ssh://git@gitlab.ost.ch:45022/ins/jalapeno-api/push-service.git
```
- Initialize okteto:
```bash
$ okteto init
```
- Replace content of okteto.yml with the following:
```yml
name: subscription-service
autocreate: true
image: okteto/golang:1
command: bash
namespace: <namespace-name>
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
  - APP_SERVER_ADDRESS=0.0.0.0:9000
  - ARANGO_DB=http://10.20.1.24:30852
  - ARANGO_DB_USER=root
  - ARANGO_DB_PASSWORD=jalapeno
  - ARANGO_DB_NAME=jalapeno
  - KAFKA_ADDRESS=10.20.1.24:30092
  - LSNODE_KAFKA_TOPIC=gobmp.parsed.ls_node_events
  - LSLINK_KAFKA_TOPIC=gobmp.parsed.ls_link_events
```
