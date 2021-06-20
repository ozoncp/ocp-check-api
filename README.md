# ocp-check-api

## Version: version not set

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

### Models

#### apiCreateTestRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| taskID | string (uint64) |  | No |
| input | string |  | No |
| output | string |  | No |

#### apiCreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test_id | string (uint64) |  | No |

#### apiDescribeTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| test | [apiTest](#apitest) |  | No |

#### apiListTestsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tests | [ [apiTest](#apitest) ] |  | No |

#### apiMultiCreateTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created | string (uint64) |  | No |

#### apiRemoveTestResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| deleted | boolean |  | No |

#### apiTest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| taskID | string (uint64) |  | No |
| input | string |  | No |
| output | string |  | No |

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
| error | string |  | No |
| code | integer |  | No |
| message | string |  | No |
| details | [ [protobufAny](#protobufany) ] |  | No |
