# Фаза 1 — Основы gRPC

## Что изучить

- [Введение в gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Основные понятия, архитектура и жизненный цикл](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [Жизненный цикл RPC](https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle)

## Что такое gRPC?

gRPC — это open-source-фреймворк для удалённого вызова процедур (Remote Procedure Call, RPC). Клиент
может вызвать метод сервера, работающего на другой машине, так, будто это локальный метод. Контракт
сервиса определяет, какие методы доступны, а также типы запросов и ответов, которые они используют.

gRPC удобен для взаимодействия между сервисами, особенно если клиенты и серверы написаны на разных
языках. Он предоставляет сгенерированные идиоматичные API клиента и сервера и поддерживает unary- и
streaming-вызовы.

## Когда предпочесть gRPC вместо REST?

Выбирайте gRPC, когда сервисам нужны строго типизированный контракт, сгенерированные клиенты и
эффективное взаимодействие между сервисами. gRPC также естественно подходит для API, в котором
используется потоковая передача данных.

REST/JSON часто лучше подходит для публичных HTTP API, простых интеграций и непосредственных
клиентов в браузере. Эти подходы могут сосуществовать: выбор зависит от потребителей и требований
API, а не от необходимости заменить один подход другим.

## Какие существуют 4 типа RPC?

- **Unary RPC**: клиент отправляет один запрос и получает один ответ.
- **Server-streaming RPC**: клиент отправляет один запрос и получает упорядоченный поток ответов.
- **Client-streaming RPC**: клиент отправляет упорядоченный поток запросов и получает один ответ.
- **Bidirectional-streaming RPC**: обе стороны независимо отправляют упорядоченные потоки сообщений.

## Основные понятия

- **Service** определяет удалённый API в файле `.proto`.
- **Method** — это RPC-операция внутри сервиса с типами сообщений запроса и ответа.
- **Message** — типизированная структура данных с именованными полями; она используется для запросов
  и ответов.
- Protocol Buffers — используемый по умолчанию в gRPC язык описания интерфейса (Interface Definition
  Language, IDL) и формат сообщений. Компилятор `protoc` и gRPC-плагины генерируют из контракта
  `.proto` типы protobuf-сообщений, а также API клиента и сервера для Go.

## Unary- и streaming-вызовы

Unary RPC имеет один запрос и один ответ, поэтому он ближе всего к обычному вызову функции.
Streaming RPC позволяет клиенту, серверу или обеим сторонам обмениваться несколькими сообщениями.
gRPC сохраняет порядок сообщений внутри каждого отдельного потока.

## Как выполняется вызов

1. Клиент вызывает сгенерированный stub (в Go он называется client), передавая сообщение запроса и,
   при необходимости, deadline.
2. gRPC отправляет вызов через HTTP/2. Каждый RPC соответствует потоку HTTP/2, а сериализованные
   Protocol Buffers передаются в кадрах данных HTTP/2.
3. Сервер получает имя метода и metadata запроса, декодирует запрос и запускает соответствующую
   реализацию метода сервиса.
4. Сервер отправляет сообщение ответа, финальный status code и необязательное status message, а
   также необязательную trailing metadata. До ответа он может отправить initial metadata.
5. При статусе `OK` клиент получает ответ и вызов завершается. Любая сторона может отменить вызов;
   deadline клиента может завершить его со статусом `DEADLINE_EXCEEDED`. Изменения, выполненные до
   отмены, автоматически не откатываются.

# Фаза 2 — gRPC в Go

## Что изучить

- [Quick start | Go](https://grpc.io/docs/languages/go/quickstart/) - [Учебник по основам |
Go](https://grpc.io/docs/languages/go/basics/) - [Справочник generated-кода |
Go](https://grpc.io/docs/languages/go/generated-code/)

## Toolchain и генерация кода

Для Go-проекта с gRPC нужны Go, компилятор Protocol Buffers (`protoc`) и два Go-плагина:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

Файл `.proto` является контрактом сервиса. Его `package` определяет типы Protocol Buffers, а `option
go_package` задаёт путь импорта Go и имя пакета для generated-кода.

Из директории `grpc-playground` весь protobuf-код playground можно перегенерировать командой:

```bash
make protos
```

Текущая цель `protos` запускает `protos-helloworld`. При добавлении нового примера для него нужно
создать отдельную цель и добавить её как зависимость `protos` в верхнеуровневом `Makefile`.

Текущая цель разворачивается в:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  helloworld/helloworld/helloworld.proto
```

`paths=source_relative` сохраняет generated-файлы рядом с исходным `.proto`. Команда создаёт или
обновляет:

- `helloworld/helloworld/helloworld.pb.go`: Go-типы protobuf-сообщений и код для их сериализации;
- `helloworld/helloworld/helloworld_grpc.pb.go`: сгенерированные API клиента и сервера gRPC.

Не редактируйте generated-файлы `*.pb.go` вручную. Измените контракт `.proto`, а затем снова
запустите `protoc`.

## От контракта к работающему приложению

Текущий пример работает по следующей схеме:

1. `helloworld.proto` объявляет сервис `Greeter`, его unary-методы `SayHello` и `SayHelloAgain`, а
   также сообщения `HelloRequest` и `HelloReply`.
2. `protoc` генерирует Go-типы сообщений, `GreeterClient`, `GreeterServer`, `NewGreeterClient` и
   `RegisterGreeterServer`.
3. Написанный вручную сервер реализует сгенерированный интерфейс `GreeterServer`.
4. Написанный вручную клиент создаёт сгенерированный `GreeterClient` и вызывает его методы.

Для unary-метода сгенерированный интерфейс сервера имеет такой вид:

```go
Method(context.Context, *Request) (*Response, error)
```

Метод сгенерированного клиента имеет такой вид:

```go
Method(ctx context.Context, request *Request, opts ...grpc.CallOption) (*Response, error)
```

## Структура сервера

Реализация сервера в `helloworld/greeter_server/main.go`:

1. встраивает `pb.UnimplementedGreeterServer` для обратной совместимости;
2. реализует методы сервиса;
3. открывает TCP listener через `net.Listen`;
4. создаёт сервер через `grpc.NewServer()`;
5. регистрирует реализацию через `pb.RegisterGreeterServer`;
6. блокируется в `Serve`, принимая соединения и передавая RPC соответствующим обработчикам.

## Структура клиента

Клиент в `helloworld/greeter_client/main.go`:

1. создаёт соединение через `grpc.NewClient` и закрывает его с помощью `defer`;
2. получает сгенерированный stub через `pb.NewGreeterClient(conn)`;
3. создаёт context с deadline в одну секунду;
4. вызывает `SayHello` и `SayHelloAgain` с `HelloRequest`;
5. читает возвращённый `HelloReply` или обрабатывает ошибку.

`insecure.NewCredentials()` подходит только для этого локального учебного примера. В реальном
сервисе используйте транспортную безопасность и подходящие credentials.

## Обзор generated streaming API

В новых generated streaming API для Go используются generics. Клиентские RPC-вызовы и серверные
RPC-обработчики безопасно запускать в конкурентных goroutine. Однако внутри одного stream нельзя
выполнять конкурентные чтения или конкурентные записи; одно чтение и одна запись могут выполняться
одновременно. Реализация streaming относится к Phase 4.

# Phase 3A — Правила совместимости

Рассматривайте `.proto` как долгоживущий API-контракт.

- Номера полей входят в wire format. Нельзя повторно использовать номер для другого смысла.
- Добавление нового поля с новым номером обычно обратно совместимо: старые клиенты его игнорируют.
- Добавление нового RPC, message или значения enum обычно безопасно, если не меняется смысл существующих элементов.
- Изменение типа или номера поля, его смысла, package или формы существующего RPC может сломать
  совместимость.
- При удалении поля зарезервируйте его номер и имя:

```proto
message Product {
  reserved 2;
  reserved "old_name";
  string id = 1;
}
```

Не путайте переименование в исходном коде с wire-compatible изменением: важны номер поля и его
сериализуемый смысл. Подробная практика `v1`/`v2` относится к Phase 3A; Buf-проверки начинаются в
Phase 3B.

# Карта Phase 3

- **Phase 3A — проектирование protobuf-контракта:** изучить `.proto`, совместимость, `v1`/`v2` и
  generated Go API, используя только `protoc` и Makefile.
- **Phase 3B — Buf contract workflow:** добавить format, lint, breaking checks, воспроизводимую
  генерацию и основы зависимостей/версий.
- **Phase 3C — инструменты вокруг protobuf/gRPC:** сравнить ConnectRPC, опубликовать тот же контракт
  через grpc-gateway/OpenAPI, диагностировать grpcurl и добавить OpenTelemetry.

Подробная последовательность находится в [roadmap gRPC +
микросервисов](../docs/learning/grpc-microservices-roadmap.ru.md). Основной путь остаётся `grpc-go +
protobuf + Buf`; Phase 3C расширяет его, но не заменяет typed Go tests и основную protobuf-механику.
