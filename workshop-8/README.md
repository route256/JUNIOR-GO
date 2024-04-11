# gRPC

## Установка protoc

https://grpc.io/docs/protoc-installation/

* Linux

```bash
$ apt install -y protobuf-compiler
$ protoc --version
```

* Mac

```bash
$ brew install protobuf
$ protoc --version 
```

* Precomplied binaries

https://github.com/protocolbuffers/protobuf/releases

## Установка плагинов

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## (опционально) buf

https://github.com/bufbuild/buf

## jaeger в Docker

```bash
$ docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

# Домашнее задание

## Задание

Требования:
1) Переписать REST сервисы из домашнего задания 6 на gRPC 
2) Переписать логирование на структурное (с исользованием go.uber.org/zap)
3) Добавить трейсы

## Дополнительно
💎 Подключить gRPC-Gateway для обратной совместимости с REST