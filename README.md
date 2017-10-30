# webhulk

[![Build Status](https://travis-ci.org/grvcoelho/webhulk.svg?branch=master)](https://travis-ci.org/grvcoelho/webhulk)

:construction: A lightweight API for managing webhooks

## API

### Webhooks

| Attribute | Type | Description |
| --------- | ---- | ----------- |
| `name` | *String* | An indentifier for the webhook |
| `url` | *String* | The url messages will be sent to |
| `enabled` | *Boolean* | Whether or not it is enabled |


### Messages

| Attribute | Type | Description |
| --------- | ---- | ----------- |
| `headers` | *JSON* | The HTTP headers that will be sent |
| `payload` | *JSON* | The payload that will be sent |
| `signature` | *String* | An HMAC signature of the message |

### Deliveries

| Attribute | Type | Description |
| --------- | ---- | ----------- |
| `status` | *String* | The status of the delivery. One of: `success`, `failed`, `processing` |
| `latency` | *JSON* | The time between the delivery and the response from the client |
| `response_status_code` | *String* | The HTTP status code the client respond with |
| `response_headers` | *JSON* | The HTTP headers the client respond with |
