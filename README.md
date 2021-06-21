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
| 200 | A successful response. | [apiListChecksResponse](#apilistchecksresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### POST
##### Summary

Обновляет "проверку" по идентификатору

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [apiUpdateCheckResponse](#apiupdatecheckresponse) |
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
| 200 | A successful response. | [apiDescribeCheckResponse](#apidescribecheckresponse) |
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
| 200 | A successful response. | [apiRemoveCheckResponse](#apiremovecheckresponse) |
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
| 200 | A successful response. | [apiListTestsResponse](#apilisttestsresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

#### POST
##### Summary

Обновляет "тест" по идентификатору

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [apiUpdateTestResponse](#apiupdatetestresponse) |
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
| 200 | A successful response. | [apiDescribeTestResponse](#apidescribetestresponse) |
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
| 200 | A successful response. | [apiRemoveTestResponse](#apiremovetestresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### /version

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [apiApiVersionResponse](#apiapiversionresponse) |
| default | An unexpected error response. | [runtimeError](#runtimeerror) |

### Models

#### apiApiVersionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| buildDateTime | string |  | No |
| gitCommit | string |  | No |
| protocolRevision | string |  | No |

#### apiCheck

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| runnerID | string (uint64) |  | No |
| solutionID | string (uint64) |  | No |
| success | boolean |  | No |
| testID | string (uint64) |  | No |

#### apiCreateCheckRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| runnerID | string (uint64) |  | No |
| solutionID | string (uint64) |  | No |
| success | boolean |  | No |
| testID | string (uint64) |  | No |

#### apiCreateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| check_id | string (uint64) |  | No |

#### apiCreateTestRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| input | string |  | No |
| output | string |  | No |
| taskID | string (uint64) |  | No |

#### apiCreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test_id | string (uint64) |  | No |

#### apiDescribeCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| check | [apiCheck](#apicheck) |  | No |

#### apiDescribeTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test | [apiTest](#apitest) |  | No |

#### apiListChecksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| checks | [ [apiCheck](#apicheck) ] |  | No |

#### apiListTestsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tests | [ [apiTest](#apitest) ] |  | No |

#### apiMultiCreateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created | string (uint64) |  | No |

#### apiMultiCreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created | string (uint64) |  | No |

#### apiRemoveCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| deleted | boolean |  | No |

#### apiRemoveTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| deleted | boolean |  | No |

#### apiTest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| input | string |  | No |
| output | string |  | No |
| taskID | string (uint64) |  | No |

#### apiUpdateCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| updated | boolean |  | No |

#### apiUpdateTestResponse

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

