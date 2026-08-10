# gRPC + микросервисы на Go: roadmap

Практический roadmap для изучения **gRPC на Go** и последующего перехода к
**микросервисному мышлению** без слабых книг и случайных фрагментов учебников.

Этот roadmap построен на идее, что сначала нужно изучить **gRPC как технологию
взаимодействия**, затем **gRPC на Go**, потом **production-механику**, и только
после этого рассматривать gRPC как часть более широкой микросервисной
архитектуры.

---

## Карта обучения

```mermaid
flowchart TD
    A[Что такое gRPC] --> B[Основные понятия и жизненный цикл RPC]
    B --> C[Go Quick Start]
    C --> D[Учебник по основам Go]
    D --> E[Проектирование Proto и generated-код]
    E --> F[Все четыре типа RPC]
    F --> G[Metadata Interceptors Auth]
    G --> H[Deadlines Cancellation Status Codes Errors]
    H --> I[Health Checking Reflection Graceful Shutdown]
    I --> J[Retry Wait-for-Ready Service Config]
    J --> K[Name Resolution Load Balancing]
    K --> L[Observability Performance Keepalive]
    L --> M[Границы микросервисов дизайн API стратегия отказов]
    M --> N[Дополнительно: Compression Flow Control Hedging gRPC-Web ALTS]
```

---

## Главный принцип

Не начинайте с «микросервисов» как с модного buzzword.

Начните в таком порядке:

1. **Основы gRPC**
2. **gRPC на Go**
3. **Production-механика**
4. **Надёжность и observability**
5. **Микросервисная архитектура вокруг gRPC**

Этот порядок важен. Иначе можно получить три сервиса, пять контейнеров и не
понять, почему запросы завершаются ошибкой. Это довольно распространённое
человеческое увлечение.

---

## Текущий прогресс

- **Фаза 1 — Основы gRPC:** завершена. Конспект находится в
  [`grpc-playground/cheatsheet.md`](../../grpc-playground/cheatsheet.md).
- **Фаза 2 — gRPC на Go:** завершена. В `grpc-playground` есть unary gRPC-сервер
  и клиент, воспроизводимая генерация protobuf-кода и integration test в памяти.
- **Текущий фокус:** Фаза 3 — Protocol Buffers и проектирование контрактов.

---

## Фаза 1. Основы gRPC

### Цель

Понять, что такое gRPC, как выполняются RPC-вызовы и чем gRPC отличается от
обычных HTTP+JSON API.

### Что изучить

- [Введение в gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Основные понятия, архитектура и жизненный цикл](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [Жизненный цикл RPC](https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle)

### Что нужно понять

- что такое **service**, **method** и **message**;
- разницу между **unary** и **streaming** RPC;
- как клиент и сервер взаимодействуют через HTTP/2;
- зачем нужны Protocol Buffers;
- что происходит во время одного RPC-вызова от запроса до ответа.

### Результат

Напишите короткую личную заметку с ответами на вопросы:

- Что такое gRPC?
- Когда я предпочту его REST?
- Какие существуют 4 типа RPC?

---

## Фаза 2. gRPC на Go

### Цель

Научиться создавать и запускать базовые gRPC-сервер и клиент на Go.

### Что изучить

- [Quick start | Go](https://grpc.io/docs/languages/go/quickstart/)
- [Учебник по основам | Go](https://grpc.io/docs/languages/go/basics/)
- [Справочник generated-кода | Go](https://grpc.io/docs/languages/go/generated-code/)

### Что нужно понять

- как определить файл `.proto`;
- как сгенерировать Go-код из `.proto`;
- как выглядят stubs сервера и клиента;
- как зарегистрировать реализацию сервиса;
- как вызвать RPC из Go-клиента.

### Результат

Создайте небольшой проект, содержащий:

- 1 gRPC-сервер;
- 1 Go-клиент;
- 1 unary-метод;
- аккуратную структуру проекта.

---

## Фаза 3. Protocol Buffers и проектирование контрактов

### Цель

Научиться воспринимать `.proto` как API-контракт, а не просто как файл с
синтаксисом, который передаётся в `protoc`.

### Что изучить

- [Protocol Buffers](https://protobuf.dev/)
- [Справочник generated-кода | Go](https://grpc.io/docs/languages/go/generated-code/)

### Что нужно понять

- `package` и `go_package`;
- messages и enums;
- номера полей и совместимость;
- различие между `optional` и `repeated`;
- соглашения об именовании;
- как generated-интерфейсы отображаются на Go-код.

### Практика

Спроектируйте две версии одного API:

- `v1` с базовым request/response;
- `v2` с обратно совместимыми дополнениями.

### Результат

Зафиксируйте для себя в документации:

- какие изменения безопасны;
- какие изменения ломают совместимость;
- как регенерация меняет Go-код.

---

## Фаза 4. Все 4 типа RPC

### Цель

Перестать воспринимать gRPC как «REST, только с компиляцией».

### Что изучить

Используйте учебник по основам Go и расширьте его собственными примерами.

### Реализовать на практике

- **Unary RPC**;
- **Server-streaming RPC**;
- **Client-streaming RPC**;
- **Bidirectional-streaming RPC**.

### Результат

Один сервис, предоставляющий все 4 типа методов.

Возможные предметные области:

- уведомления;
- логи/события;
- отслеживание маршрутов;
- демонстрация в стиле чата;
- сборщик метрик.

---

## Фаза 5. Metadata, interceptors и auth

### Цель

Понять сквозное поведение в gRPC.

### Что изучить

- [Metadata](https://grpc.io/docs/guides/metadata/)
- [Interceptors](https://grpc.io/docs/guides/interceptors/)
- [Authentication](https://grpc.io/docs/guides/auth/)

### Что нужно понять

- metadata запроса и ответа;
- unary- и stream-interceptors;
- передачу auth-токена;
- передачу tracing-данных или request ID;
- простые шаблоны middleware в gRPC.

### Практика

Добавьте на сервер:

- request ID в metadata;
- logging interceptor;
- auth-check interceptor;
- interceptor для измерения времени и latency.

### Результат

Переиспользуемый пакет interceptors для учебного проекта.

---

## Фаза 6. Deadlines, cancellation, status codes и errors

### Цель

Корректно обрабатывать ошибки вместо случайного `fmt.Errorf` с привкусом
грусти.

### Что изучить

- [Deadlines](https://grpc.io/docs/guides/deadlines/)
- [Cancellation](https://grpc.io/docs/guides/cancellation/)
- [Status codes](https://grpc.io/docs/guides/status-codes/)
- [Error handling](https://grpc.io/docs/guides/error/)

### Что нужно понять

- как клиенты устанавливают deadlines;
- как серверы учитывают отмену через context;
- канонические status codes;
- отображение бизнес-ошибок на gRPC-ошибки;
- почему timeout является частью поведения API.

### Практика

Реализуйте сценарии:

- некорректный ввод;
- объект не найден;
- доступ запрещён;
- зависимость недоступна;
- deadline exceeded.

### Результат

Небольшая таблица отображения ошибок для проекта.

---

## Фаза 7. Основы эксплуатации

### Цель

Сделать сервис пригодным для работы в реальном окружении.

### Что изучить

- [Health checking](https://grpc.io/docs/guides/health-checking/)
- [Server reflection](https://grpc.io/docs/guides/reflection/)
- [Graceful shutdown](https://grpc.io/docs/guides/server-graceful-stop/)

### Что нужно понять

- как работают health probes;
- почему reflection полезен при разработке и отладке;
- как остановить сервер, не прерывая выполняющиеся запросы.

### Практика

Добавьте:

- health service;
- reflection в режиме разработки;
- обработку сигналов с graceful shutdown.

### Результат

Сервис, который можно запустить, проверить и корректно остановить.

---

## Фаза 8. Надёжность на стороне клиента

### Цель

Понять, как gRPC-клиенты ведут себя при сбоях.

### Что изучить

- [Retry](https://grpc.io/docs/guides/retry/)
- [Wait-for-Ready](https://grpc.io/docs/guides/wait-for-ready/)
- [Service Config](https://grpc.io/docs/guides/service-config/)
- [Request Hedging](https://grpc.io/docs/guides/request-hedging/)

### Что нужно понять

- transparent retry и настроенный retry;
- политики для отдельных методов;
- ограничения retry;
- как wait-for-ready меняет поведение при ошибках;
- почему бездумные retry могут усиливать сбои.

### Практика

Создайте нестабильный downstream-сервис и проверьте:

- немедленную ошибку;
- retry policy;
- взаимодействие retry с deadline;
- поведение wait-for-ready.

### Результат

Короткая заметка о том, когда retry безопасен, а когда опасен.

---

## Фаза 9. Name resolution и load balancing

### Цель

Понять, как клиенты находят серверы и распределяют вызовы.

### Что изучить

- [Custom name resolution](https://grpc.io/docs/guides/custom-name-resolution/)
- [Load balancing](https://grpc.io/docs/guides/custom-load-balancing/)
- [Service Config](https://grpc.io/docs/guides/service-config/)

### Что нужно понять

- роли resolver и balancer;
- статическое и динамическое обнаружение targets;
- основы round robin;
- как service config влияет на поведение клиента.

### Практика

Запустите 2 экземпляра одного сервиса и проверьте стратегию балансировки.

### Результат

Мини-демо с одним клиентом и несколькими backend-инстансами.

---

## Фаза 10. Observability и производительность

### Цель

Измерять и настраивать поведение вместо того, чтобы гадать.

### Что изучить

- [OpenTelemetry Metrics](https://grpc.io/docs/guides/opentelemetry-metrics/)
- [Performance Best Practices](https://grpc.io/docs/guides/performance/)
- [Keepalive](https://grpc.io/docs/guides/keepalive/)
- [Flow Control](https://grpc.io/docs/guides/flow-control/)
- [Compression](https://grpc.io/docs/guides/compression/)

### Что нужно понять

- количество запросов, latency и ошибки;
- метрики на стороне сервера и клиента;
- повторное использование соединений;
- компромиссы keepalive;
- backpressure и flow control в streaming;
- когда compression помогает, а когда мешает.

### Практика

Добавьте метрики и сравните:

- пропускную способность unary и streaming;
- маленькие и большие payloads;
- работу с compression и без неё;
- работу с настройкой keepalive и без неё.

### Результат

Простой benchmark report в Markdown.

---

## Фаза 11. Микросервисы вокруг gRPC

### Цель

Использовать gRPC в осмысленной архитектуре из нескольких сервисов.

### Важное замечание

Официальная документация gRPC хорошо объясняет **механику RPC**, но **не
является полным курсом по микросервисной архитектуре**. Эту часть нужно
изучать отдельно.

### Рекомендуемые дополнительные roadmaps

- [Go roadmap](https://roadmap.sh/golang)
- [Backend roadmap](https://roadmap.sh/backend)
- [API Design roadmap](https://roadmap.sh/api-design)
- [System Design roadmap](https://roadmap.sh/system-design)

### Что изучить

- границы сервисов;
- синхронное и асинхронное взаимодействие;
- распространение отказов;
- идемпотентность;
- контракты между сервисами;
- observability за границами сервисов;
- когда *не* следует выделять отдельный сервис.

### Практический проект

Создайте небольшую систему из 3 сервисов:

- `users`;
- `orders`;
- `payments`.

Добавьте:

- gRPC-вызовы между сервисами;
- health checks;
- deadlines;
- retries;
- структурированное логирование;
- метрики;
- симуляцию одной нестабильной зависимости.

### Результат

README с описанием:

- архитектуры;
- ответственности сервисов;
- критических сценариев отказа;
- правил retry и timeout.

---

## Фаза 12. Дополнительные advanced topics

Изучайте их после основного пути, а не раньше.

### Темы

- [gRPC-Web](https://grpc.io/docs/platforms/web/)
- [ALTS](https://grpc.io/docs/guides/alts/)
- [Debugging](https://grpc.io/docs/guides/debugging/)
- [Custom backend metrics](https://grpc.io/docs/guides/custom-backend-metrics/)
- [Reflection](https://grpc.io/docs/guides/reflection/) в сценариях с большим количеством инструментов

### Примечания

- **gRPC-Web** важен, если клиентом является браузер.
- **ALTS** не относится к первоочередным темам для обычной backend-разработки.
- **Инструменты отладки** становятся полезнее после появления реальных сервисов.

---

# Рекомендуемый порядок обучения: Сейчас / Далее / Позже

## Сейчас

- [x] Введение в gRPC
- [x] Основные понятия
- [x] Go Quick Start
- [x] Учебник по основам Go
- [x] Справочник generated-кода
- [x] Один unary RPC на Go
- [ ] Все 4 типа RPC в одном демонстрационном сервисе

## Далее

- [ ] Проектирование Proto и совместимость
- [ ] Metadata
- [ ] Interceptors
- [ ] Основы аутентификации
- [ ] Deadlines
- [ ] Cancellation
- [ ] Status codes
- [ ] Error handling
- [ ] Health checking
- [ ] Reflection
- [ ] Graceful shutdown

## Позже

- [ ] Retry
- [ ] Wait-for-ready
- [ ] Service Config
- [ ] Name resolution
- [ ] Load balancing
- [ ] Метрики OpenTelemetry
- [ ] Настройка производительности
- [ ] Keepalive
- [ ] Flow control
- [ ] Compression
- [ ] Учебный multi-service проект
- [ ] gRPC-Web / ALTS / дополнительные advanced topics

---

# Минимальная последовательность практических проектов

## Проект 1. `grpc-playground`

Один Go-сервер и один Go-клиент.

Должны присутствовать:

- unary RPC;
- server streaming RPC;
- client streaming RPC;
- bidirectional streaming RPC;
- простой `.proto`.

## Проект 2. `grpc-runtime`

Превратите playground в более реалистичный сервис.

Добавьте:

- metadata;
- interceptors;
- проверку auth-токена;
- deadlines;
- обработку cancellation;
- status codes;
- graceful shutdown;
- health checking;
- reflection.

## Проект 3. `grpc-micro-lab`

Три сервиса с реалистичным поведением.

Добавьте:

- gRPC-вызовы между сервисами;
- retry policies;
- симуляцию нестабильной зависимости;
- метрики;
- правила timeout;
- один эксперимент с load balancing;
- архитектурный README.

---

# Как пользоваться этим roadmap

Используйте roadmap последовательно:

1. Прочитайте официальную страницу.
2. Соберите минимальный работающий пример.
3. Намеренно сломайте его.
4. Запишите, что произошло на самом деле.
5. Переходите к следующей теме только после практики.

В этом и состоит весь секрет. Не гламурно, не мистично, зато эффективно.

---

# Напутствие

Не ищите один магический ресурс, который идеально обучит **gRPC + Go +
микросервисам + production + архитектуре + observability**.

Такого ресурса почти наверняка не существует.

Лучше:

- использовать **официальную документацию gRPC** как основной каркас;
- использовать **практические проекты на Go** для наработки навыков;
- использовать материалы по **backend и system design** как архитектурный слой.

Так вы получите настоящий путь обучения, а не конфетти из случайных tutorial.
