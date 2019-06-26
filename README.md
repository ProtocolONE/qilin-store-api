[![Build Status](https://travis-ci.org/ProtocolONE/qilin-store-api.svg?branch=master)](https://travis-ci.org/ProtocolONE/qilin-store-api.svg)

# Qilin Store API

Qilin Store Api is an open source backend service for Cord project.

## Get started

Qilin Store API designed to be launched with Kubernetes and handle all configuration from env variables:

| Variable                               | Default      | Description                                                                                                                                |
|----------------------------------------|--------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| QILINSTOREAPI_SERVER_PORT              | 8080                 | HTTP port to listed API requests.                                                                                                          |
| QILINSTOREAPI_SERVER_ALLOW_ORIGINS     | *                    | Comma separated list of [CORS domains](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin).             |
| QILINSTOREAPI_SERVER_ALLOW_CREDENTIALS | false                | Look at [CORS documentation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) about this value. |
| QILINSTOREAPI_CACHE_HOST               | debug                | Default logging level in application.                                                                                                      |
| QILINSTOREAPI_DB_HOST                  | 127.0.0.1            | Mongo DB host.                                                                                                      |
| QILINSTOREAPI_DB_NAME                  | qilinstoreapi        | Mongo DB database name.                                                                                                      |
| QILINSTOREAPI_DB_USER                  |                      | Mongo DB user.                                                                                                      |
| QILINSTOREAPI_DB_PASSWORD              |                      | Mongo DB password.                                                                                                      |
| QILINSTOREAPI_DB_EVENT_BUS             |amqp://127.0.0.1:5672 | RabbitMQ event bus connection string.                                                                                                      |
| QILINSTOREAPI_SESSION_STORAGE_HOST     |localhost             | Redis host for session storage                                                                                                     |
| QILINSTOREAPI_SESSION_STORAGE_PORT     |6379                  | Redis port for session storage                                                                                                     |
| QILINSTOREAPI_SESSION_STORAGE_PASSWORD |                      | Redis password for session storage                                                                                                     |
| QILINSTOREAPI_SESSION_STORAGE_SECRET   |                      | Secret key for sessions encryption                                                                                                     |
| QILINSTOREAPI_CACHE_HOST               | localhost            | Redis host  for caching                                                                                               |
| QILINSTOREAPI_CACHE_PORT               | 6379                 | Redis port  for caching                                                                                               |
| QILINSTOREAPI_CACHE_PASSWORD           |                      | Redis password for caching                                                                                               |
| QILINSTOREAPI_AUTH1_ISSUER             |                      | URL to ProtocolOne authentication server (without slash on the end).                                                                       |
| QILINSTOREAPI_AUTH1_CLIENTID           |                      | Application identifier from ProtocolOne authenticate server.                                                                               |
| QILINSTOREAPI_AUTH1_CLIENTSECRET       |                      | Secret authentication key for the application from ProtocolOne authenticate server.                                                        |

## Supported go versions
We support the major Go versions, which are 1.11 at the moment.

## Contributing
Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue. Extra credit for those using Testify to write the test code that demonstrates it.
