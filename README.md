# ocp-check-api

API для работы с сущностями check/test (проверка/тест) в составе Ozon Code Platform. Представляет собой микросервис, принимающий запросы через gRPC и поддерживающий операции Create/Read/Update/Delete с записью в БД PostgreSQL и отправку в Kafka.

## Сборка

Сборка осуществляется командой 

```bash
$ make build
```

Исполняемый файл будет размещен в директории bin.

## Запуск

Перед запуском создайте в текущей директории или в директории config файл config.yml со следующими параметрами

```yaml
database:
  # url для подключения к БД. Параметр дублируется системной переменной ocp_check_api_database_url
  url: postgres://postgres:postgres@localhost:5432/postgres
grpc:
  # интерфейс и порт gRPC-сервера, который "поднимается" микросервисом. Параметр дублируется системной переменной ocp_check_api_grpc_address
  address: 0.0.0.0:8083
kafka:
  # адреса и порты инстансов Kafka, в которые будут отправляться сообщения по операциям Create/Update/Delete. Параметр дублируется системной переменной ocp_check_api_kafka_brokers
  brokers: [127.0.0.1:9092, 127.0.0.1:9093]
prometheus:
  # интерфейс и порт Prometheus-метрик, которые "выдаются" микросервисом. Параметр дублируется системной переменной ocp_check_api_prometheus_address
  address: 0.0.0.0:9100
log:
  # Вывод в консоль. Можно включить, установив значение enable: true
  console:
    enable: false
  file:
   # Вывод в лог. Можно отключить, установив значение enable: false. Значение path определяет путь к логу.
    enable: true
    path: /var/log/ocp-check-api.log
```

Запуск микросервиса осуществляется командой
```bash
$ ./bin/main
```

## Миграция БД

Файлы для миграций БД размещаются с директории migrations.
Перед запуском миграции необходимо скачать утилиту migrate. Скачивание осуществляется командой 

```bash
$ make migrate
```

В корневом каталоге репозитория должен появиться исполняемый файл migrate, для миграции нужно запустить его с параметрами
```bash
$ ./migrate -source file:./migrations -database postgres://localhost:5432/db?sslmode=require up 2
```

# Protocol description
## Version: 1

### /checks

#### GET
##### Summary

Возвращает список "проверок"

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string (uint64) |
| offset | query |  | No | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListChecksResponse](#listchecksresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### POST
##### Summary

Обновляет "проверку" по идентификатору

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateCheckResponse](#updatecheckresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### /checks/{check_id}

#### GET
##### Summary

Возвращает описание "проверки" по ее идентификатору

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| check_id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DescribeCheckResponse](#describecheckresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### DELETE
##### Summary

Удаляет "проверку" по идентификатору

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| check_id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [RemoveCheckResponse](#removecheckresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### /tests

#### GET
##### Summary

Возвращает список "тестов"

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string (uint64) |
| offset | query |  | No | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListTestsResponse](#listtestsresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### POST
##### Summary

Обновляет "тест" по идентификатору

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateTestResponse](#updatetestresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### /tests/{test_id}

#### GET
##### Summary

Возвращает описание "теста" по ее идентификатору

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| test_id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DescribeTestResponse](#describetestresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### DELETE
##### Summary

Удаляет "тест" по идентификатору

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| test_id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [RemoveTestResponse](#removetestresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### /version

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ApiVersionResponse](#apiversionresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### Models

#### ApiVersionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| buildDateTime | string |  | No |
| gitCommit | string |  | No |
| protocolRevision | string |  | No |

#### Check

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| runnerID | string (uint64) |  | No |
| solutionID | string (uint64) |  | No |
| success | boolean |  | No |
| testID | string (uint64) |  | No |

#### CreateCheckRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| runnerID | string (uint64) |  | No |
| solutionID | string (uint64) |  | No |
| success | boolean |  | No |
| testID | string (uint64) |  | No |

#### CreateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| check_id | string (uint64) |  | No |

#### CreateTestRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| input | string |  | No |
| output | string |  | No |
| taskID | string (uint64) |  | No |

#### CreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test_id | string (uint64) |  | No |

#### DescribeCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| check | [Check](#check) |  | No |

#### DescribeTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test | [Test](#test) |  | No |

#### ListChecksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| checks | [ [Check](#check) ] |  | No |

#### ListTestsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tests | [ [Test](#test) ] |  | No |

#### MultiCreateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created | string (uint64) |  | No |

#### MultiCreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created | string (uint64) |  | No |

#### RemoveCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| deleted | boolean |  | No |

#### RemoveTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| deleted | boolean |  | No |

#### Test

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| input | string |  | No |
| output | string |  | No |
| taskID | string (uint64) |  | No |

#### UpdateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| updated | boolean |  | No |

#### UpdateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| updated | boolean |  | No |

#### protobufAny

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| type_url | string |  | No |
| value | byte |  | No |

#### runtimeError

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| details | [ [protobufAny](#protobufany) ] |  | No |
| error | string |  | No |
| message | string |  | No |
